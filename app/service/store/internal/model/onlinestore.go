package model

import "encoding/json"

// 线上门店
type OnlineStore struct {
	Store
	StoreCount
	OnlineStoreAttr
	OnlineStoreAddress
}

type OnlineStoreAttr struct {
	CommonStoreAttr
}

type OnlineStoreAddress struct {
	CommonStoreAddress
}

func (o *OnlineStore) CloneAddress(s string){
	address := &OnlineStoreAddress{}
	_ = address.Decode(s)

	o.ProvinceName = address.ProvinceName
	o.CityName = address.CityName
	o.AreaName = address.AreaName
	o.ProvinceCode = address.ProvinceCode
	o.CityCode = address.CityCode
	o.AreaCode = address.AreaCode
	o.Address = address.Address
	o.Gcj02 = address.Gcj02
	o.GaodeId = address.GaodeId
	o.Location = address.Location
	o.Wechat = address.Wechat
	o.Tel = address.Tel
}

func (o *OnlineStore) CloneStore(s *Store){
	o.ID = s.ID
	o.Title = s.Title
	o.Type = s.Type
	o.Level = s.Level
	o.Cover = s.Cover
	o.Source = s.Source
	o.Displayorder = s.Displayorder
	o.Status = s.Status
	o.StatusTime = s.StatusTime
	o.TagIds = s.TagIds

	return
}

func (o *OnlineStore) CloneAttr(s string){
	oa := &OnlineStoreAttr{}
	oa.Decode(s)

	o.BusinessTime = oa.BusinessTime
	o.Remark = oa.Remark
	o.Styles = oa.Styles
	o.Introduction = oa.Introduction
	o.PriceDesc = oa.PriceDesc
	o.BrandDesc = oa.BrandDesc
	o.ProductsDesc = oa.ProductsDesc
	o.BusinessTime = oa.BusinessTime
	return
}

func (o *OnlineStoreAttr) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), o); err != nil{
		return
	}
	return
}

func (o *OnlineStoreAttr) Encode() (s string, err error){
	var b []byte
	if b,err = json.Marshal(o); err != nil{
		return "", nil
	}
	s = string(b)
	return
}

func (o *OnlineStoreAddress) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), o); err != nil{
		return
	}
	return
}

func (o *OnlineStoreAddress) Encode() (s string, err error) {
	var b []byte
	if b,err = json.Marshal(o); err != nil{
		return "", nil
	}
	return string(b), nil
}

















