package model

type Pagination struct {
	RecordCount int
	PageCount int
	Current int
	PageSize int
}

type Location struct {
	Lat float64
	Lon float64
}

type StoreIdsByEsResult struct {
	pagination *Pagination
	StoreIds []int
}

type StoreEsSearchParams struct {
	Order string
	Type int
	BrandId int
	DealerId int
	Tag string
	Source int
	City string
	Distance string
	Location *Location
	Page int
	PageSize int
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


