package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/google/wire"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/dao"
	"go-common/app/service/store/internal/model"
)

var Provider = wire.NewSet(New)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Close close the resource.
func (s *Service) Close() {

}

func (s *Service) AddStore(ctx context.Context, req *v1.EditStoreReq) (storeId int, err error) {
	storeId, err = s.dao.AddStore(ctx, req)
	log.Info("result:%v", err)
	return
}

func (s *Service) EditStore(ctx context.Context, req *v1.EditStoreReq) (err error) {
	err = s.dao.EditStore(ctx, req)
	log.Info("result:%v", err)
	return
}

func (s *Service) StoreDetail(ctx context.Context, storeId int) (err error) {
	store,err := s.dao.StoreDetail(ctx, storeId)
	log.Info("result:%v", store)
	return
}

func (s *Service) StoreList(ctx context.Context, v *v1.StoreListReq) (rep *v1.StoreListRep ,err error) {
	var(
		total int
		storeIds []int
	)

	req := &model.StoreEsSearchParams{
		BrandId:           int(v.BrandId),
		DealerId:          int(v.DealerId),
		PageSize:          int(v.PageSize),
		Page:              int(v.Page),
	}
	if total,storeIds,err = s.dao.StoreEsSearch(ctx, req); err != nil{
		return
	}

	stores := make([]*v1.Store, 0, len(storeIds))
	storeList,_ := s.dao.StoreByIds(ctx, storeIds)

	for _, value := range storeList{
		_store := &v1.Store{
			Id:                   int32(value.ID),
			Title:                value.Title,
			Type:                 int32(value.Type),
			Level:                int32(value.Level),
			Cover:                value.Cover,
		}
		stores = append(stores, _store)
	}
	rep = &v1.StoreListRep{
		Store:                stores,
		Page:&v1.Page{
			Page:                 v.Page,
			PageSize:             v.PageSize,
			RecordCount:          int32(total),
		},
	}
	return
}

func (s *Service) StoreFuzzySearch(ctx context.Context) (err error) {
	req := &v1.FuzzySearchStoreReq{
		Query:                "索菲亚",
		Order:                0,
		Page:                 1,
		PageSize:             20,
	}
	total,storeIds,err := s.dao.StoreFuzzySearch(ctx, req)
	//storeList,err := s.dao.StoreByIds(ctx, storeIds)
	log.Info("total:%s  ids:%v  err:%v", total, storeIds, err)
	//for _, value := range storeList{
	//	log.Info("store data: %v", value.Title)
	//}
	return
}

func (s *Service) Test(ctx context.Context, storeId int) (err error) {
	// 线下门店详情
	//list,err := s.dao.OnlineStoreDetail(ctx, storeId)
	//log.Info("result:%v", err)
	//log.Info("result:%v", list)

	// 加图片资源
	//id,err := s.dao.AddStoreResourceImage(ctx, storeId, "图片1", "/test/11/bo.jpg", &model.StoreResourceImageExt{
	//	W: 100,
	//	H: 200,
	//})
	//log.Info("result:%v", err)
	//log.Info("result:%v", id)

	// 加视频资源
	//id,err := s.dao.AddStoreResourceVideo(ctx, storeId, "图片1", "/test/11/bo.mp4", &model.StoreResourceVideoExt{
	//	Mime:  "mp4",
	//	W:     200,
	//	H:     300,
	//	Cover: model.StoreResourceVideoExtCover{
	//		W:   200,
	//		H:   300,
	//		Url: "/test/11/bo.mp4.jpg",
	//	},
	//})
	//log.Info("result:%v", err)
	//log.Info("result:%v", id)

	// 图片资源列表
	//list,err := s.dao.StoreResourceImageList(ctx, storeId)
	//for _, val := range list{
	//	log.Info("data i11s %v", val.Title)
	//}
	//log.Info("result:%v", err)
	//log.Info("result:%v", list)

	// 视频资源列表
	//list,err := s.dao.StoreResourceVideoList(ctx, storeId)
	//for _, val := range list{
	//	log.Info("data i11s %v", val.Title)
	//}
	//log.Info("result:%v", err)
	//log.Info("result:%v", list)

	// 增加门店从属关系
	//err = s.dao.AddStoreBelone(ctx, storeId, 2, "5,6,8")
	//log.Info("result:%v", err)

	// 根据ids获取门店
	list,_ := s.dao.StoreByIds(ctx, []int{15,12,13})
	for _,val := range list{
		log.Info("title is : %v", val.Title)
	}

	// 测试修改count表
	//err = s.dao.SetStoreCount(ctx, 19, "images", 10)
	//log.Info("result:%v", err)


	// es搜索
	//req := &model.StoreEsSearchParams{
	//	Order:             "",
	//	Type:              0,
	//	Keywords:          "",
	//	BrandId:           0,
	//	DealerId:          0,
	//	Tag:               "",
	//	Source:            0,
	//	City:              "",
	//	InsertStoreIds:    "",
	//	AreaCode:          0,
	//	Distance:          "",
	//	DistrictLocations: nil,
	//	Location:          nil,
	//}
	//_,_ = s.dao.OfflineStoreEsSearch(ctx, req)
	return
}




