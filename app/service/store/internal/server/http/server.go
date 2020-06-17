package http

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	_ "github.com/go-kratos/kratos/pkg/net/rpc/warden"
	_ "go-common/app/service/store/api"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/service"
	"strconv"
)

var svr *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svr = s
	engine = bm.DefaultServer(&cfg)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	e.GET("/mygrpc", mygrpc)
	e.GET("/addstore", addStore)
	e.GET("/editstore", editStore)
	e.GET("/test", test)
	e.GET("/detail", storeDetail)
}

func ping(ctx *bm.Context) {

}

func mygrpc(ctx *bm.Context) {
	//add := &v1.AddStoreReq{
	//	Title:                 "3232323",
	//	Type:11,
	//	Level:25,
	//	Cover:"23232",
	//	Source:11,
	//}
	//_ = svr.StoreList(ctx, add)


	//cfg := &warden.ClientConfig{}
	//offlineStoreClient, _ := v1.NewClient(cfg)
	//req := new(v1.OfflineStoreListReq)
	//var data string
	//rep,err := offlineStoreClient.OfflineStoreList(c, req) // 这个调用有我问题
	//
	//if err != nil{
	//	data = "23"
	//}else{
	//	storeId := int64(0)
	//	for _, val := range rep.OfflineStore {
	//		storeId = val.StoreId
	//	}
	//	data = fmt.Sprintf("page is %v", storeId)
	//}
	ctx.String(200, "1111")
}
func addStore(ctx *bm.Context) {
	add := &v1.AddStoreReq{
		UserId:               0,
		Title:                "3232323",
		Type:                 1,
		Cover:                "/prod/content/202006/13/836dfa/2306b7664efef14c756a641175b5fa7f.jpg",
		Images:               "/prod/content/202006/13/836dfa/2306b7664efef14c756a641175b5fa7f.jpg",
		//BelongId:             0,
		//BrandIds:             "",
		//DealerIds:            "",
		//GaodeId:              "",
		ProvinceCode:         440000,
		CityCode:             440100,
		AreaCode:             440100,
		//Source:               0,
		//TagIds:               "",
		Address:              "广东省广州市海珠区昌岗街道",
		Tel:                  "020-03121",
		//Gcj_02:               "",
		Location:             "113.43,53.12",
		//DistrictId:           0,
		Wechat:               "wx_adfdfdf",
		Introduction:         "门店描述",
		BrandNames:           "品牌名字",
		Styles:               "中国风",
		Products:             "很多产品",
		PriceDescs:           "价格丰简由人",
		BusinessTime:         "每天都营业",
	}
	storeId,_ := svr.AddStore(ctx, add)
	ctx.String(200, "add success storeId:"+strconv.FormatInt(storeId, 10))
}

func editStore(ctx *bm.Context) {
	edit := &v1.EditStoreReq{
		Title: "updatebozi",
		Type:1,
		Level:25,
		Cover:"338888",
		Source:11,
		Id:13,
		Styles:"英伦风",
		CityCode:11010,

	}
	_ = svr.EditStore(ctx, edit)
	ctx.String(200, "edit success")
}

func storeDetail(ctx *bm.Context) {
	_ = svr.StoreDetail(ctx, 5)
	ctx.String(200, "detail success")
}

func test(ctx *bm.Context) {
	storeId,_ := strconv.ParseInt(ctx.Request.Form.Get("store_id"), 10, 64)
	_ = svr.Test(ctx, storeId)
	ctx.String(200, "test success")
}

