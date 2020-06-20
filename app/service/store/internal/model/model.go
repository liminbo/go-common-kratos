package model

type Pagination struct {
	RecordCount int64
	PageCount int64
	Current int64
	PageSize int64
}

type Location struct {
	Lat string
	Lon string
}

type StoreIdsByEsResult struct {
	pagination *Pagination
	StoreIds []int64
}

type StoreEsSearchParams struct {
	Order int8
	Type int8
	Keywords string
	BrandId int64
	DealerId int64
	Tag string
	Source int8
	City string
	InsertStoreIds string
	AreaCode int64
	Distance string
	DistrictLocations []*Location
	Location *Location
}

type StoreCover struct {
	Title string
	Urls []string // 视频链接
	Cover string // 配图
	CoverInfo interface{}
}

type ImageParams struct {
	Url string `json:"url"`
	Title string `json:"title"`
	Ext StoreResourceImageExt `json:"ext"`
}

type VideoParams struct {
	Url string `json:"url"`
	Title string `json:"title"`
	Ext StoreResourceVideoExt `json:"ext"`
}


