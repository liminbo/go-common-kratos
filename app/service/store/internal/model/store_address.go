package model

type CommonStoreAddress struct {
	ProvinceName string `json:"province_name"`
	CityName string `json:"city_name"`
	AreaName string `json:"area_name"`
	ProvinceCode int `json:"province_code"`
	CityCode int `json:"city_code"`
	AreaCode int  `json:"area_code"`
	Address string  `json:"address"` // 详细地址
	Gcj02 string    `json:"gcj_02"` // 高德坐标
	GaodeId string  `json:"gaode_id"` // 高德id
	Location string `json:"location"` // 定位
	Wechat string `json:"wechat"` // 微信号
	Tel string `json:"tel"` // 电话
}
