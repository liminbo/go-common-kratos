package dao

// 责任链模式

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/model"
)

// 上下文
type StoreCtx struct {
	ctx context.Context
	storeId int64
	dao *dao
	req interface{}
	addReq *v1.AddStoreReq
	editReq *v1.EditStoreReq
}

func (ctx *StoreCtx) New(c context.Context, d *dao, addReq *v1.AddStoreReq,  editReq *v1.EditStoreReq,) (sc *StoreCtx, err error){
	sc = &StoreCtx{
		dao: d,
		ctx: c,
		addReq: addReq,
		editReq: editReq,
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

type StoreAddResourceHandle struct {
	StoreNext
}

type StoreCountHandle struct {
	StoreNext
}

type StoreEditHandle struct {
	StoreNext
}

type StoreEditResourceHandle struct {
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
	req := ctx.addReq

	store := &model.Store{
		Title:        req.Title,
		Type:         req.Type,
		Level:        req.Level,
		Cover:        req.Cover,
		Source:       req.Source,
		TagIds:       req.TagIds,
	}
	if err = ctx.dao.db.Create(store).Error; err != nil{
		log.Error("db error:%v", err)
		return
	}
	ctx.storeId = store.ID
	return
}

func (c *StoreAddResourceHandle) Do(ctx *StoreCtx) (err error){
	var id int64

	if ctx.addReq.Images != ""{
		if id,err = ctx.dao.AddStoreResourceImage(ctx.ctx, ctx.storeId, "图片title", ctx.addReq.Images, nil); err != nil{
			return
		}
	}
	log.Info("add resource id :%v", id)
	return
}

func (a *StoreAttrHandle) Do(ctx *StoreCtx) (err error){
	var attrStr string

	if ctx.addReq != nil{ // 添加
		req := ctx.addReq
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
	}else if ctx.editReq != nil{ // 编辑
		req := ctx.editReq
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
	}

	if err = ctx.dao.db.Model(&model.Store{}).Where("id=?", ctx.storeId).UpdateColumn("attr", attrStr).Error; err != nil{
		log.Error("db error: %v", err)
	}

 	return
}



func (a *StoreAddressHandle) Do(ctx *StoreCtx) (err error){
	var addressStr string

	if ctx.addReq != nil{
		req := ctx.addReq
		switch req.Type {
		case model.STORE_TYPE_OFFLINE:
			address := &model.OfflineStoreAddress{
				CommonStoreAddress: model.CommonStoreAddress{
					ProvinceName: "",
					CityName:     "",
					AreaName:     "",
					ProvinceCode: req.ProvinceCode,
					CityCode:     req.CityCode,
					AreaCode:     req.AreaCode,
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
					ProvinceCode: req.ProvinceCode,
					CityCode:     req.CityCode,
					AreaCode:     req.AreaCode,
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
	}else if ctx.editReq != nil{
		req := ctx.editReq
		switch req.Type {
		case model.STORE_TYPE_OFFLINE:
			address := &model.OfflineStoreAddress{
				CommonStoreAddress: model.CommonStoreAddress{
					ProvinceName: "",
					CityName:     "",
					AreaName:     "",
					ProvinceCode: req.ProvinceCode,
					CityCode:     req.CityCode,
					AreaCode:     req.AreaCode,
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
					ProvinceCode: req.ProvinceCode,
					CityCode:     req.CityCode,
					AreaCode:     req.AreaCode,
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

func (c *StoreEditHandle) Do(ctx *StoreCtx) (err error){
	req := ctx.req.(*v1.EditStoreReq)
	store := &model.Store{
		ID:req.Id,
		Title:        req.Title,
		Type:         req.Type,
		Level:        req.Level,
		Cover:        req.Cover,
		Source:       req.Source,
		TagIds:       req.TagIds,
	}
	if err = ctx.dao.db.Save(store).Error; err != nil{
		log.Error("mysql err %v", err)
		return
	}
	ctx.storeId = store.ID
	return
}

func (c *StoreEditResourceHandle) Do(ctx *StoreCtx) (err error){
	log.Info("store resource edit handle")
	return
}

