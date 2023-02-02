package dao

import (
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/mistake"
	"hongcha/go-expire/pkg/pager"
)

// AddProduct 添加商品
func AddProduct(product *model.Product) (err error) {
	_, err = db.Engine.Insert(product)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// GetProductAll 获取所有商品
func GetProductAll() (product []*model.Product, err error) {
	err = db.Engine.Find(&product)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// UpdateProductStatus 更新商品状态
func UpdateProductStatus(id, state int) (err error) {
	_, err = db.Engine.ID(id).Cols("state").Update(&model.Product{State: state})
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// DeleteProduct 删除商品
func DeleteProduct(id int) (err error) {
	_, err = db.Engine.ID(id).Delete(&model.Product{})
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// GetProductListPage 查询物品 分页
func GetProductListPage(req *val.GetProductWithPageReq) (product []*val.GetProductWithPageResp, total int64, err error) {
	cond := &model.Product{}
	session := db.Engine.NewSession()
	session.Table("product")
	session.Select("product.name, product.manufacture_date, product.quality_guarantee_period, product.remind, category.name as category_name, category.icon")
	if req.Name != "" {
		session.And("name like concat('%',?,'%')", req.Name)
	}
	if req.CategoryID != 0 {
		session.And("category_id = ?", req.CategoryID)
	}
	session.Join("INNER", "category", "product.category_id = category.id")
	total, err = pager.Help(req.Page, req.PageSize, &product, cond, session)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}
