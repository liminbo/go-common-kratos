package model

import (
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/log"
)

// 门店资源
const(
	STORE_RESOURCE_TYPE_IMAGE = 0 // 图片
	STORE_RESOURCE_TYPE_VIDEO = 1 // 视频
)

type StoreResourcer interface {
	FormatData() (i interface{}, e error)
}

// 资源
type StoreResource struct{
	Id int // 资源id
	StoreId int // 门店id
	Type int // 资源类型
	Title string //
	Data string // 资源数据
	DataInfo string // 资源信息
	Displayorder int // 排序
	CreatedAt int
	UpdatedAt int
}

// 图片
type StoreResourceImage struct{
	StoreResource
	StoreResourceImageExt
}
type StoreResourceImageExt struct{
	W int `json:"w"`
	H int `json:"h"`
}

// 视频
type StoreResourceVideo struct{
	StoreResource
	StoreResourceVideoExt
}
type StoreResourceVideoExt struct{
	Mime string `json:"mime"`
	W int `json:"w"`
	H int `json:"h"`
	Cover StoreResourceVideoExtCover `json:"cover"`
}

type StoreResourceVideoExtCover struct {
	W int `json:"w"`
	H int `json:"h"`
	Url string `json:"url"`
}

func (i *StoreResourceImage) FormatData() (s string, err error) {
	s = i.Data
	return
}

func (i *StoreResourceImageExt) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), i); err != nil{
		return
	}
	return
}

func (i *StoreResourceImageExt) Encode() (s string, err error) {
	var b []byte
	if b,err = json.Marshal(i); err != nil{
		return "", nil
	}
	return string(b), nil
}


func (i *StoreResourceImage) Clone(r *StoreResource){
	ext := new(StoreResourceImageExt)
	if err := ext.Decode(r.DataInfo);err != nil{
		log.Error("json decode error:%v", err)
	}

	i.Id = r.Id
	i.Type = r.Type
	i.Title = r.Title
	i.Data = r.Data
	i.DataInfo = r.DataInfo
	i.StoreId = r.StoreId
	i.Displayorder = r.Displayorder
	i.W = ext.W
	i.H = ext.H
	return
}

func (v *StoreResourceVideo) FormatData() (s string, err error) {
	s = v.Data
	return
}

func (i *StoreResourceVideo) Clone(r *StoreResource){
	ext := new(StoreResourceVideoExt)
	if err := ext.Decode(r.DataInfo);err != nil{
		log.Error("json decode error:%v", err)
	}

	i.Id = r.Id
	i.Type = r.Type
	i.Title = r.Title
	i.Data = r.Data
	i.DataInfo = r.DataInfo
	i.StoreId = r.StoreId
	i.Displayorder = r.Displayorder
	i.W = ext.W
	i.H = ext.H
	i.Cover = ext.Cover
	return
}

func (v *StoreResourceVideoExt) Decode(s string) (err error) {
	if err = json.Unmarshal([]byte(s), v); err != nil{
		return
	}
	return
}

func (v *StoreResourceVideoExt) Encode() (s string, err error) {
	var b []byte
	if b,err = json.Marshal(v); err != nil{
		return "", nil
	}
	return string(b), nil
}