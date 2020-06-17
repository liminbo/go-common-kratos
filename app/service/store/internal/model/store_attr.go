package model

// 公用属性
type CommonStoreAttr struct {
	Reply string `json:"reply"`	// 回复内容
	Remark string `json:"remark"`   // 备注
	Styles string `json:"styles"` // 风格
	Introduction string `json:"introduction"` // 介绍
	PriceDesc string `json:"price_desc"` // 价格描述
	BrandDesc string `json:"brand_desc"` // 经营品牌描述
	ProductsDesc string `json:"products_desc"` // 售卖产品描述
	BusinessTime string `json:"business_time"` // 营业时间
}
