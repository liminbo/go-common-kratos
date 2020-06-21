package dao

import (
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/log"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/model"
	elastic "gopkg.in/olivere/elastic.v5"
	"strconv"
)

import (
	"context"
)

const (
	_storeIndex   = "ydt_offlinestore_online"
	//_videoType    = "video_info"
	_videoMapping = ``

	SmartNew = "smart"

)


func (d *dao) StoreEsSearch(ctx context.Context, searchParams *model.StoreEsSearchParams) (total int, storeIds []int, err error) {
	var(
		result *elastic.SearchResult
	)
	queryBool := elastic.NewBoolQuery()
	queryBool.Filter(elastic.NewTermQuery("status", 1))


	if searchParams.BrandId != 0{
		queryBool.Must(elastic.NewTermsQuery("brand_ids", searchParams.BrandId))
	}
	if searchParams.DealerId != 0{
		queryBool.Must(elastic.NewTermsQuery("dealer_ids", searchParams.DealerId))
	}
	// 距离搜索
	if searchParams.Location != nil && searchParams.Distance != ""{
		locationQuery := elastic.NewGeoDistanceQuery("location").Distance("10km").Lat(searchParams.Location.Lat).Lon(searchParams.Location.Lon)
		queryBool.Must(locationQuery)
	}


	funcScore := elastic.NewFunctionScoreQuery()
	funcScore = funcScore.Query(queryBool)

	funcQuery1 := elastic.NewRangeQuery("related_item_num").Gte(10)
	funcScore = funcScore.Add(funcQuery1, elastic.NewWeightFactorFunction(5))

	// 排序
	sorterSlice := []elastic.Sorter{}
	sorterSlice = append(sorterSlice, elastic.NewScoreSort().Desc())
	sorterSlice = append(sorterSlice, elastic.NewFieldSort("level").Desc())


	// 打印查询数据 start
	debugSource,_ := funcScore.Source()
	debug,_ := json.Marshal(debugSource)
	log.Info("query:%s", string(debug))

	for _, value := range sorterSlice{
		_sort,_ := value.Source()
		debug,_ := json.Marshal(_sort)
		log.Info("sort:%s", string(debug))
	}
	// 打印查询数据 end

	search := d.elasticSearch.Client.Search().Index(_storeIndex)
	if result,err = search.Timeout("60s").Query(funcScore).SortBy(sorterSlice...).From(0).Size(int(searchParams.PageSize)).Do(ctx); err != nil{
		return
	}

	storeIds = make([]int, 0, searchParams.PageSize)
	total = int(result.TotalHits())
	for _,value := range result.Hits.Hits{
		tmp := new(model.StoreEsSource)
		if err = json.Unmarshal(*value.Source, tmp); err != nil{
			log.Error("search result json err %v", err)
		}
		storeId,_ := strconv.Atoi(tmp.StoreId)
		storeIds = append(storeIds, storeId)
	}
	log.Info("store ids:%v", storeIds)
	return
}

func (d *dao) StoreFuzzySearch(ctx context.Context, req *v1.FuzzySearchStoreReq) (total int, storeIds []int, err error) {
	var(
		result *elastic.SearchResult
	)
	queryBool := elastic.NewBoolQuery()
	queryBool.Must(elastic.NewTermQuery("status", 1))
	queryBool.Must(elastic.NewTermQuery("dianpin_shop_id", 0))
	queryBool.Must(elastic.NewRangeQuery("level").Gte(0))
	queryBool.Must(elastic.NewMatchPhraseQuery("name", req.Query))

	queryBool2 := elastic.NewBoolQuery()
	queryBool2.MustNot(elastic.NewTermQuery("source", 3))
	queryBool.Filter(queryBool2)

	// score评分
	scoreFuncQuery := elastic.NewFunctionScoreQuery()
	scriptStr := "double ratio = 1;if(doc['source'].value == 3){ ratio = ratio + 1 } if(doc['level'].value == 2){ ratio = ratio + 1} return ratio * _score;"
	script := elastic.NewScriptInline(scriptStr).Lang("painless")
	scoreFunc := elastic.NewScriptFunction(script)

	// 整合function_score query
	scoreFuncQuery.Query(queryBool).AddScoreFunc(scoreFunc)

	debugSource,_ := scoreFuncQuery.Source()
	debug,_ := json.Marshal(debugSource)
	log.Info("query:%s", string(debug))


	search := d.elasticSearch.Client.Search().Index(_storeIndex)
	if result,err = search.Query(scoreFuncQuery).From(0).Size(int(req.PageSize)).Do(ctx); err != nil{
		return
	}

	storeIds = make([]int, 0, req.PageSize)
	total = int(result.TotalHits())
	for _,value := range result.Hits.Hits{
		tmp := new(model.StoreEsSource)
		if err = json.Unmarshal(*value.Source, tmp); err != nil{
			log.Error("search result json err %v", err)
		}
		storeId,_ := strconv.ParseInt(tmp.StoreId, 10, 32)
		storeIds = append(storeIds, int(storeId))
	}
	return
}

func (d *dao) StoreEsSearchParams(ctx context.Context, req *v1.StoreListReq) (p *model.StoreEsSearchParams, err error) {
	search := d.elasticSearch.Client.Search().Index(_storeIndex)
	result,err := search.From(0).Size(10).Do(ctx)
	log.Info("result:%v  err:%v", result, err)
	return
}
