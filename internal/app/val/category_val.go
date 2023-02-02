package val

// GetCategoryResp 分类 返回结构
type GetCategoryResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetCategoryReq 分类 请求结构体
type GetCategoryReq struct {
	// 分类名称
	Name string `json:"name"`
	// 分类图标📈
	Icon string `json:"icon"`
}
