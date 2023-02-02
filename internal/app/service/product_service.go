package service

import (
	"github.com/jinzhu/copier"
	"hongcha/go-expire/internal/app/dao"
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/mistake"
	"hongcha/go-expire/pkg/pager"
	"time"
)

// AddProduct 添加商品
func AddProduct(req *val.GetProductReq) (resp string, err error) {

	if req.Name == "" {
		err = mistake.NewReqErr(mistake.ErrUnknown, "商品名称不能为空")
		return
	}

	if req.ManufactureDate == "" {
		err = mistake.NewReqErr(mistake.ErrUnknown, "生产日期不能为空")
		return
	}

	if req.QualityGuaranteePeriod == 0 {
		err = mistake.NewReqErr(mistake.ErrUnknown, "保质期不能为空")
		return
	}

	if req.CategoryID == 0 {
		err = mistake.NewReqErr(mistake.ErrUnknown, "商品分类不能为空")
		return
	}

	if req.Unit == "" {
		err = mistake.NewReqErr(mistake.ErrUnknown, "单位不能为空")
		return
	}

	qualityGuaranteePeriod := time.Now().AddDate(0, 0, req.QualityGuaranteePeriod)

	// 字符串时间转换为时间
	manufactureDate, _ := time.Parse("2006-01-02", req.ManufactureDate)

	if req.Unit == "day" {
		qualityGuaranteePeriod = manufactureDate.AddDate(0, 0, req.QualityGuaranteePeriod)
	}

	if req.Unit == "month" {
		qualityGuaranteePeriod = manufactureDate.AddDate(0, req.QualityGuaranteePeriod, 0)
	}

	if req.Unit == "year" {
		qualityGuaranteePeriod = manufactureDate.AddDate(req.QualityGuaranteePeriod, 0, -1)
	}

	product := model.Product{
		Name:                   req.Name,
		ManufactureDate:        manufactureDate,
		QualityGuaranteePeriod: qualityGuaranteePeriod,
		State:                  1,
		CategoryID:             req.CategoryID,
		Remind:                 req.Remind,
	}

	err = dao.AddProduct(&product)
	if err != nil {
		return "", mistake.New500ServiceErr(mistake.ErrUnknown, err)
	}
	return "", nil
}

// GetProductListPage 查询物品 分页
func GetProductListPage(req *val.GetProductWithPageReq) (pageModel *pager.PageModel, err error) {
	product := &model.Product{}
	_ = copier.Copy(product, req)

	authorList, total, err := dao.GetProductListPage(req)

	if err != nil {
		return nil, mistake.New500ServiceErr(mistake.ErrUnknown, err)
	}

	resp := make([]*val.GetProductWithPageResp, 0)

	for _, author := range authorList {
		t := &val.GetProductWithPageResp{}
		_ = copier.Copy(t, author)
		resp = append(resp, t)
	}

	return pager.NewPageModel(req.Page, req.PageSize, total, resp), nil
}
