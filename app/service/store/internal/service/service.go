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

func (s *Service) StoreList(ctx context.Context, req *v1.StoreListReq) (rep *v1.StoreListRep ,err error) {
	var(
		total int
		storeIds []int
	)

	params := &model.StoreEsSearchParams{
		BrandId:           int(req.BrandId),
		DealerId:          int(req.DealerId),
		PageSize:          int(req.PageSize),
		Page:              int(req.Page),
	}
	if total,storeIds,err = s.dao.StoreEsSearch(ctx, params); err != nil{
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
			Page:                 req.Page,
			PageSize:             req.PageSize,
			RecordCount:          int32(total),
		},
	}
	return
}

func (s *Service) StoreFuzzySearch(ctx context.Context, req *v1.FuzzySearchStoreReq) (rep *v1.FuzzySearchStoreRep, err error) {
	total,storeIds,err := s.dao.StoreFuzzySearch(ctx, req)
	log.Info("fuzzy search storeIds:%v", storeIds)

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
	rep = &v1.FuzzySearchStoreRep{
		Store:                stores,
		Page:&v1.Page{
			Page:                 req.Page,
			PageSize:             req.PageSize,
			RecordCount:          int32(total),
		},
	}

	return
}

func (s *Service) DeleteStore(ctx context.Context, storeId int) (err error) {
	err = s.dao.DeleteStore(ctx, storeId)
	return
}

func (s *Service) StoreImageList(ctx context.Context, storeId int) (rep *v1.StoreImageListRep, err error) {
	rep = new(v1.StoreImageListRep)

	list,err := s.dao.StoreResourceImageList(ctx, storeId)
	images := make([]*v1.ResourceImage, 0, len(list))

	for _,val := range list{
		image := &v1.ResourceImage{
			Title:                val.Title,
			Data:                 val.Data,
			Ext:                  &v1.ResourceImage_Ext{
				W:                    int32(val.StoreResourceImageExt.W),
				H:                    int32(val.StoreResourceImageExt.H),
			},
		}
		images = append(images, image)
	}

	rep.Image = images
	return
}

func (s *Service) StoreVideoList(ctx context.Context, storeId int) (rep *v1.StoreVideoListRep, err error) {
	rep = new(v1.StoreVideoListRep)

	list,err := s.dao.StoreResourceVideoList(ctx, storeId)
	videoes := make([]*v1.ResourceVideo, 0, len(list))

	for _,val := range list{
		video := &v1.ResourceVideo{
			Title:                val.Title,
			Data:                 val.Data,
			Ext:                  &v1.ResourceVideo_Ext{
				W:                    int32(val.W),
				H:                    int32(val.H),
				Mime:                 val.Mime,
				Cover:                &v1.ResourceVideo_Cover{
					W:                    int32(val.Cover.W),
					H:                    int32(val.Cover.H),
					Url:                  val.Cover.Url,
				},
			},
		}
		videoes = append(videoes, video)
	}

	rep.Video = videoes
	return
}

func (s *Service) OfflineStoreDetail(ctx context.Context, storeId int) (rep *v1.OfflineStoreDetailRep, err error) {
	var(
		detail *model.OfflineStore
	)
	rep = new(v1.OfflineStoreDetailRep)

	if detail,err = s.dao.OfflineStoreDetail(ctx, storeId);err != nil{
		return
	}
	offlineStore := &v1.OfflineStore{
		Id:                   int32(detail.ID),
		Title:                detail.Title,
		Type:                 int32(detail.Type),
		Level:                int32(detail.Level),
		Cover:                detail.Cover,
		ProvinceName:         detail.ProvinceName,
		CityName:             detail.CityName,
		AreaName:             detail.AreaName,
		ProvinceCode:         int32(detail.ProvinceCode),
		CityCode:             int32(detail.CityCode),
		AreaCode:             int32(detail.AreaCode),
		Address:              detail.Address,
		Tel:                  detail.Tel,
		Gcj_02:               detail.Gcj02,
		Location:             detail.Location,
		DistrictId:           int32(detail.DistrictId),
		Wechat:               detail.Wechat,
		Reply:                detail.Reply,
		Remark:               detail.Remark,
		Styles:               detail.Styles,
		Introduction:         detail.Introduction,
		PriceDesc:            detail.PriceDesc,
		BrandDesc:            detail.BrandDesc,
		ProductsDesc:         detail.ProductsDesc,
		BusinessTime:         detail.BusinessTime,
		Articles:             int32(detail.Articles),
		Evaluations:          int32(detail.Evaluations),
		Items:                int32(detail.Items),
		RealClicks:           int32(detail.RealClicks),
		Clicks:               int32(detail.Clicks),
		Images:               int32(detail.Images),
		Videoes:              int32(detail.Videoes),
		CreatedAt:            0,
	}

	rep.OfflineStore = offlineStore
	return
}

func (s *Service) OnlineStoreDetail(ctx context.Context, storeId int) (rep *v1.OnlineStoreDetailRep, err error) {
	var(
		detail *model.OnlineStore
	)
	rep = new(v1.OnlineStoreDetailRep)

	if detail,err = s.dao.OnlineStoreDetail(ctx, storeId);err != nil{
		return
	}
	onlineStore := &v1.OnlineStore{
		Id:                   int32(detail.ID),
		Title:                detail.Title,
		Type:                 int32(detail.Type),
		Level:                int32(detail.Level),
		Cover:                detail.Cover,
		ProvinceName:         detail.ProvinceName,
		CityName:             detail.CityName,
		AreaName:             detail.AreaName,
		ProvinceCode:         int32(detail.ProvinceCode),
		CityCode:             int32(detail.CityCode),
		AreaCode:             int32(detail.AreaCode),
		Address:              detail.Address,
		Tel:                  detail.Tel,
		Gcj_02:               detail.Gcj02,
		Location:             detail.Location,
		Wechat:               detail.Wechat,
		Reply:                detail.Reply,
		Remark:               detail.Remark,
		Styles:               detail.Styles,
		Introduction:         detail.Introduction,
		PriceDesc:            detail.PriceDesc,
		BrandDesc:            detail.BrandDesc,
		ProductsDesc:         detail.ProductsDesc,
		BusinessTime:         detail.BusinessTime,
		Articles:             int32(detail.Articles),
		Evaluations:          int32(detail.Evaluations),
		Items:                int32(detail.Items),
		RealClicks:           int32(detail.RealClicks),
		Clicks:               int32(detail.Clicks),
		Images:               int32(detail.Images),
		Videoes:              int32(detail.Videoes),
		CreatedAt:            0,
	}

	rep.OnlineStore = onlineStore
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




