package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/google/wire"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/dao"
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

func (s *Service) AddStore(ctx context.Context, req *v1.AddStoreReq) (storeId int64, err error) {
	storeId, err = s.dao.AddStore(ctx, req)
	log.Info("result:%v", err)
	return
}

func (s *Service) EditStore(ctx context.Context, req *v1.EditStoreReq) (err error) {
	err = s.dao.EditStore(ctx, req)
	log.Info("result:%v", err)
	return
}

func (s *Service) StoreDetail(ctx context.Context, storeId int64) (err error) {
	store,err := s.dao.StoreDetail(ctx, storeId)
	log.Info("result:%v", store)
	return
}

func (s *Service) Test(ctx context.Context, storeId int64) (err error) {
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
	//list,_ := s.dao.StoreByIds(ctx, []int64{12,13,15})
	//for _,val := range list{
	//	log.Info("title is : %v", val.Title)
	//}


	return
}




