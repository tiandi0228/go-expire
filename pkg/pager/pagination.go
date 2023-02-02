package pager

import "reflect"

// PageModel 分页模型
type PageModel struct {
	Page         int         `json:"page"`
	PageSize     int         `json:"page_size"`
	TotalPages   int64       `json:"total_pages"`
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
}

// NewPageModel 构造分页对象
// page: 当前页码
// pageSize: 每页条数
// totalRecords: 总条数
// records: 分页数据，必须为slice，不能为nil或者其他对象，可以是空的slice
func NewPageModel(page, pageSize int, totalRecords int64, records interface{}) *PageModel {
	sliceValue := reflect.Indirect(reflect.ValueOf(records))
	if sliceValue.Kind() != reflect.Slice {
		panic("分页异常，需要传入一个slice类型")
	}

	page, pageSize = ValPageAndPageSize(page, pageSize)
	if totalRecords < 0 {
		totalRecords = 0
	}

	totalPages := CalcTotalPages(int64(pageSize), totalRecords)

	return &PageModel{
		Page:         page,
		PageSize:     pageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		Records:      records,
	}
}

// CalcTotalPages 计算页码逻辑
func CalcTotalPages(pageSize, totalRecords int64) (totalPages int64) {
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages
}

// ValPageAndPageSize 保证分页参数合法性
func ValPageAndPageSize(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}

// CalcStartEndPage 计算分页参数，start和end，用于excel分页预览
func CalcStartEndPage(page, pageSize, totalRecords int) (start, end int) {
	page, pageSize = ValPageAndPageSize(page, pageSize)
	if totalRecords < 0 {
		totalRecords = 0
	}

	start = (page - 1) * pageSize
	end = start + pageSize
	if end > totalRecords {
		end = totalRecords
	}
	return
}
