package model


type Storer interface {

}
const(
	// 门店上下架
	STORE_STATUS_HIDE = 0  // 下架
	STORE_STATUS_SHOW = 1  // 上架

	// 门店等级
	STORE_LEVEL_LOW = 0
	STORE_LEVEL_MIDDLE = 1
	STORE_LEVEL_HIGH = 2

	// 门店来源
	STORE_SOURCE_ADMIN = 0 // 后台
	STORE_SOURCE_SCRIPT = 1 // 脚本
	STORE_SOURCE_DIANPIN = 2 // 点评抓取
	STORE_SOURCE__COLLECTION = 3 // 线下采集

	//门店类型
	STORE_TYPE_OFFLINE = 1
	STORE_TYPE_ONLINE = 2
)
type Store struct {
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` // 门店id
	Title string // 名字
	Type int // 类型
	Level int // 等级
	Cover string // 封面图
	Source int // 来源
	Displayorder int // 上下架状态
	Status int // 审核状态
	StatusTime int // 审核时间
	TagIds string // 标签id集
	Attr string // 相关属性
	Address string // 地址

	CreatedAt int
	UpdatedAt int
}

// 门店商区商圈
type StoreDistrict struct {
	Id int
	Name string // 商区名
	CityCode int
	AreaCode int
	parentcid int // 父ID
	IsParent int  // 是否父商区
	IsHot int	// 是否热门
	StoreNum int // 门店数
	Location string // 定位
	CreatedAt int
	UpdatedAt int
}

// 门店从属表
const(
	STORE_BELONG_TYPE_INDEPENDENT = 0 // 独立
	STORE_BELONG_TYPE_BRAND = 1 //品牌
	STORE_BELONG_TYPE_DEALER = 2 //经销商
)
type StoreBelong struct{
	StoreId int // 门店id
	BelongType int // 从属类型
	BelongIds string // 从属id集
	CreatedAt int
	UpdatedAt int
}

// 门店统计表
type StoreCount struct {
	StoreId int // 门店id
	Articles int // 文章数
	Evaluations int // 评价数
	Items int // 商品数
	RealClicks int // 真实点击数
	Clicks int // 点击数
	Images int // 图片数
	Videoes int // 视频数
	CreatedAt int
	UpdatedAt int

}

type StoreEsSource struct {
	StoreId string `json:"store_id"`
}