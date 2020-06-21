package dao

// 责任链模式

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/log"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/model"
)

// 上下文
type StoreCtx struct {
	ctx context.Context
	storeId int
	dao *dao
	req *v1.EditStoreReq
}

func (ctx *StoreCtx) New(c context.Context, d *dao, editReq *v1.EditStoreReq,) (sc *StoreCtx, err error){
	sc = &StoreCtx{
		dao: d,
		ctx: c,
		req: editReq,
	}
	return
}

type StoreHandler interface {
	Do(ctx *StoreCtx) error
	setNext(h StoreHandler) StoreHandler
	Run(ctx *StoreCtx) error
}

type StoreNext struct {
	NextHandle StoreHandler
}

type StoreNilHandle struct {
	StoreNext
}

type StoreAddHandle struct {
	StoreNext
}

type StoreAttrHandle struct {
	StoreNext
}

type StoreAddressHandle struct {
	StoreNext
}

type StoreResourceHandle struct {
	StoreNext
}

type StoreCountHandle struct {
	StoreNext
}

type StoreBelongHandle struct {
	StoreNext
}

type StoreEditHandle struct {
	StoreNext
}

func (n *StoreNext) setNext(h StoreHandler) StoreHandler{
	n.NextHandle = h
	return h
}

func (n *StoreNext) Run(ctx *StoreCtx) (err error){
	if n.NextHandle != nil{
		err = (n.NextHandle).Do(ctx)
		if err != nil{
			return
		}
		return (n.NextHandle).Run(ctx)
	}
	return
}

func (n *StoreNilHandle) Do(ctx *StoreCtx) (err error){
	return
}

func (a *StoreAddHandle) Do(ctx *StoreCtx) (err error){
	req := ctx.req

	store := &model.Store{
		Title:        req.Title,
		Type:         int(req.Type),
		Level:        int(req.Level),
		Cover:        req.Cover,
		Source:       int(req.Source),
		TagIds:       req.TagIds,
	}
	if err = ctx.dao.db.Create(store).Error; err != nil{
		log.Error("db error:%v", err)
		return
	}
	ctx.storeId = store.ID
	return
}

func (c *StoreResourceHandle) Do(ctx *StoreCtx) (err error){

	var(
		id int
		images []model.ImageParams
		videos []model.VideoParams
		imageCount int = 0
		videoCount int = 0
	)

	// 先删除资源记录,再插入
	if err = ctx.dao.DeleteStoreResource(ctx.ctx, ctx.storeId); err != nil{
		return
	}

	// 处理图片 start
	if err = json.Unmarshal([]byte(ctx.req.Images), &images); err != nil{
		log.Error("format image err: %v", err)
	}
	if images != nil{
		imageCount = len(images)
		for _, image := range images{
			if id,err = ctx.dao.AddStoreResourceImage(ctx.ctx, ctx.storeId, image.Title, image.Url, &image.Ext); err != nil{
				return
			}
		}
	}
	_ = ctx.dao.SetStoreCount(ctx.ctx, ctx.storeId, "images", imageCount)
	// 处理图片 end

	// 处理视频 start
	if err = json.Unmarshal([]byte(ctx.req.Videos), &videos); err != nil{
		log.Error("format Video err: %v", err)
	}
	if videos != nil{
		videoCount = len(videos)
		for _, video := range videos{
			if id,err = ctx.dao.AddStoreResourceVideo(ctx.ctx, ctx.storeId, video.Title, video.Url, &video.Ext); err != nil{
				return
			}
		}
	}
	_ = ctx.dao.SetStoreCount(ctx.ctx, ctx.storeId, "videoes", videoCount)
	// 处理视频 end

	log.Info("add resource id :%v", id)
	return
}

func (a *StoreAttrHandle) Do(ctx *StoreCtx) (err error){
	var attrStr string

	req := ctx.req
	switch req.Type {
		case model.STORE_TYPE_OFFLINE:
			attr := &model.OfflineStoreAttr{
				CommonStoreAttr: model.CommonStoreAttr{
					Reply:        "",
					Remark:       "",
					Styles:       req.Styles,
					Introduction: req.Introduction,
					PriceDesc:    req.PriceDescs,
					BrandDesc:    req.BrandNames,
					ProductsDesc: req.Products,
					BusinessTime: req.BusinessTime,
				},
			}
			attrStr, _ = attr.Encode()
			// 格式化属性 end

		case model.STORE_TYPE_ONLINE:
			attr := &model.OnlineStoreAttr{
				CommonStoreAttr: model.CommonStoreAttr{
					Reply:        "",
					Remark:       "",
					Styles:       req.Styles,
					Introduction: req.Introduction,
					PriceDesc:    req.PriceDescs,
					BrandDesc:    req.BrandNames,
					ProductsDesc: req.Products,
					BusinessTime: req.BusinessTime,
				},
			}
			attrStr, _ = attr.Encode()
	}

	if err = ctx.dao.db.Model(&model.Store{}).Where("id=?", ctx.storeId).UpdateColumn("attr", attrStr).Error; err != nil{
		log.Error("db error: %v", err)
	}

 	return
}



func (a *StoreAddressHandle) Do(ctx *StoreCtx) (err error){
	var addressStr string

	req := ctx.req
	switch req.Type {
		case model.STORE_TYPE_OFFLINE:
			address := &model.OfflineStoreAddress{
				CommonStoreAddress: model.CommonStoreAddress{
					ProvinceName: "",
					CityName:     "",
					AreaName:     "",
					ProvinceCode: int(req.ProvinceCode),
					CityCode:     int(req.CityCode),
					AreaCode:     int(req.AreaCode),
					Address:      req.Address,
					Gcj02:        req.Gcj_02,
					GaodeId:      req.GaodeId,
					Location:     req.Location,
					Wechat:       req.Wechat,
					Tel:          req.Tel,
				},
				DistrictId:         0,
			}
			addressStr,_ = address.Encode()
		case model.STORE_TYPE_ONLINE:
			address := &model.OfflineStoreAddress{
				CommonStoreAddress: model.CommonStoreAddress{
					ProvinceName: "",
					CityName:     "",
					AreaName:     "",
					ProvinceCode: int(req.ProvinceCode),
					CityCode:     int(req.CityCode),
					AreaCode:     int(req.AreaCode),
					Address:      req.Address,
					Gcj02:        req.Gcj_02,
					GaodeId:      req.GaodeId,
					Location:     req.Location,
					Wechat:       req.Wechat,
					Tel:          req.Tel,
				},
				DistrictId:         0,
			}
			addressStr,_ = address.Encode()
	}

	if err = ctx.dao.db.Model(&model.Store{}).Where("id=?", ctx.storeId).UpdateColumn("address", addressStr).Error; err != nil{
		log.Error("db error: %v", err)
	}
	return
}

func (c *StoreCountHandle) Do(ctx *StoreCtx) (err error){
	count := &model.StoreCount{
		StoreId:     ctx.storeId,
	}
	if err = ctx.dao.db.Create(count).Error; err != nil{
		log.Error("db error:%v", err)
	}
	return
}

func (c *StoreBelongHandle) Do(ctx *StoreCtx) (err error){
	if err = ctx.dao.DeleteStoreBelong(ctx.ctx, ctx.storeId); err != nil{
		return
	}
	if ctx.req.BrandIds != "" {
		if err = ctx.dao.AddStoreBelone(ctx.ctx, ctx.storeId, model.STORE_BELONG_TYPE_BRAND, ctx.req.BrandIds); err != nil{
			return
		}
	}
	if ctx.req.DealerIds != "" {
		if err = ctx.dao.AddStoreBelone(ctx.ctx, ctx.storeId, model.STORE_BELONG_TYPE_DEALER, ctx.req.DealerIds); err != nil{
			return
		}
	}
	return
}

func (c *StoreEditHandle) Do(ctx *StoreCtx) (err error){
	req := ctx.req
	store := &model.Store{
		ID:           int(req.Id),
		Title:        req.Title,
		Type:         int(req.Type),
		Level:        int(req.Level),
		Cover:        req.Cover,
		Source:       int(req.Source),
		TagIds:       req.TagIds,
	}
	if err = ctx.dao.db.Save(store).Error; err != nil{
		log.Error("mysql err %v", err)
		return
	}
	ctx.storeId = store.ID
	return
}

