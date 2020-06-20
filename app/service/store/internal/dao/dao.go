package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/model"
	"go-common/library/database/es"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// 获取es搜索参数
	StoreEsSearchParams(ctx context.Context, req *v1.StoreListReq) (*model.StoreEsSearchParams, error)
	// 格式化门店
	StoreByIds(ctx context.Context, storeIds []int64) ([]*model.Store, error)
	// es条件搜索门店
	StoreEsSearch(ctx context.Context, searchParams *model.StoreEsSearchParams) (*model.StoreIdsByEsResult, error)
	// es全文搜索
	StoreFuzzySearch(ctx context.Context, req *v1.FuzzySearchStoreReq) (*model.StoreIdsByEsResult, error)

	// 门店详情
	StoreDetail(ctx context.Context, storeId int64) (*model.Store, error)

	// 线下门店详情
	OfflineStoreDetail(ctx context.Context, storeId int64) (*model.OfflineStore, error)

	// 线上门店详情
	OnlineStoreDetail(ctx context.Context, storeId int64) (*model.OnlineStore, error)

	// 添加门店
	AddStore(ctx context.Context, req *v1.EditStoreReq) (int64, error)
	// 编辑门店
	EditStore(ctx context.Context, req *v1.EditStoreReq) error
	// 删除门店
	DeleteStore(ctx context.Context, storeId int64) error

	// 获取资源列表
	StoreResourceList(ctx context.Context, storeId,resourceType int64) ([]*model.StoreResource, error)
	// 设置门店资源
	AddStoreResource(ctx context.Context, resource *model.StoreResource) (int64,error)

	AddStoreResourceImage(ctx context.Context, storeId int64, title, data string, dataInfo *model.StoreResourceImageExt) (id int64, err error)

	AddStoreResourceVideo(ctx context.Context, storeId int64, title, data string, dataInfo *model.StoreResourceVideoExt) (id int64, err error)

	StoreResourceImageList(ctx context.Context, storeId int64) (list []*model.StoreResourceImage, err error)
	StoreResourceVideoList(ctx context.Context, storeId int64) (list []*model.StoreResourceVideo, err error)

	// 设置门店统计
	SetStoreCount(ctx context.Context, storeId int64, field string, count int) error
	// 获取门店统计
	StoreCount(ctx context.Context, storeId int64) (*model.StoreCount, error)
	AddStoreBelone(ctx context.Context, storeId int64, belongType int8, belongIds string) (err error)
	// 获取门店统计
	StoreBelone(ctx context.Context, storeId int64) ([]*model.StoreBelong, error)
	// 获取门店地区
	StoreDistrictList(ctx context.Context, storeId int64) ([]*model.StoreDistrict, error)

}

// dao dao.
type dao struct {
	db          *gorm.DB
	redis       *redis.Redis
	elasticSearch *es.ElasticSearch
	cache *fanout.Fanout
	demoExpire int32
}

// New new a dao and return.
func New(r *redis.Redis, db *gorm.DB, elasticSearch *es.ElasticSearch) (d Dao, cf func(), err error) {
	return newDao(r, db, elasticSearch)
}

func newDao(r *redis.Redis, db *gorm.DB, elasticSearch *es.ElasticSearch) (d *dao, cf func(), err error) {
	var cfg struct{
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db: db,
		redis: r,
		elasticSearch: elasticSearch,
		cache: fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	 d.cache.Close()
	 d.db.Close()
	 d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
