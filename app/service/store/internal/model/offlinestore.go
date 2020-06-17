package model

import "encoding/json"

// 线下门店
type OfflineStore struct {
	Store
	StoreCount
	OfflineStoreAttr
	OfflineStoreAddress
}

type OfflineStoreAddress struct {
	CommonStoreAddress
	DistrictId int64 // 商区id
}

type OfflineStoreAttr struct {
	CommonStoreAttr
}

func (o *OfflineStore) CloneStore(s *Store){
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


func (o *OfflineStore) CloneAddress(s string){
	address := &OfflineStoreAddress{}
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
	o.GaodeId = address.GaodeId
	o.Location = address.Location
	o.Wechat = address.Wechat
	o.Tel = address.Tel
	o.DistrictId = address.DistrictId
}

func (o *OfflineStore) CloneAttr(s string){
	oa := &OfflineStoreAttr{}
	oa.Decode(s)

	o.BusinessTime = oa.BusinessTime
	o.Remark = oa.Remark
	o.Styles = oa.Styles
	o.Introduction = oa.Introduction
	o.PriceDesc = oa.PriceDesc
	o.BrandDesc = oa.BrandDesc
	o.ProductsDesc = oa.ProductsDesc
	o.BusinessTime = oa.BusinessTime
}

func (o *OfflineStoreAttr) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), o); err != nil{
		return
	}
	return
}

func (o *OfflineStoreAttr) Encode() (s string, err error){
	var b []byte
	if b,err = json.Marshal(o); err != nil{
		return "", nil
	}
	s = string(b)
	return
}

func (o *OfflineStoreAddress) Encode() (s string, err error) {
	var b []byte
	if b,err = json.Marshal(o); err != nil{
		return "", nil
	}
	return string(b), nil
}

func (o *OfflineStoreAddress) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), o); err != nil{
		return
	}
	return
}
