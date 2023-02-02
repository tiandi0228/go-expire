package val

import "time"

// GetProductReq 物品 请求结构体
type GetProductReq struct {
	// 名称
	Name string `json:"name"`
	// 生产日期
	ManufactureDate string `json:"manufacture_date"`
	// 保质期
	QualityGuaranteePeriod int `json:"quality_guarantee_period"`
	// 分类
	CategoryID int `json:"category_id"`
	// 单位
	Unit string `json:"unit"`
	// 提醒时间
	Remind int `json:"remind"`
}

// GetProductWithPageReq 查询物品列表 分页 请求结构体
type GetProductWithPageReq struct {
	// 物品名称
	Name string `comment:"物品名称" form:"name"`
	// 分类
	CategoryID int `comment:"分类id" form:"category_id"`
	// 页码
	Page int `validate:"omitempty,min=1" comment:"页码" form:"page"`
	// 每页大小
	PageSize int `validate:"omitempty,min=1" comment:"每页大小" form:"page_size"`
}

// GetProductWithPageResp 物品 返回结构体
type GetProductWithPageResp struct {
	// 名称
	Name string `json:"name"`
	// 生产日期
	ManufactureDate time.Time `json:"manufacture_date"`
	// 保质期
	QualityGuaranteePeriod time.Time `json:"quality_guarantee_period"`
	// 分类名称
	CategoryName string `json:"category_name"`
	// 提醒时间
	Remind int `json:"remind"`
	// 分类icon
	Icon string `json:"icon"`
}
