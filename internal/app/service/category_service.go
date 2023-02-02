package service

import (
	"hongcha/go-expire/internal/app/dao"
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/mistake"

	"github.com/jinzhu/copier"
)

// GetCategoryAll 获取所有分类
func GetCategoryAll() (resp []*val.GetCategoryResp, err error) {
	categoryAll, err := dao.GetCategoryAll()
	if err != nil {
		return nil, mistake.New500ServiceErr(mistake.ErrUnknown, err)
	}

	resp = make([]*val.GetCategoryResp, 0)

	_ = copier.Copy(&resp, &categoryAll)
	return
}

// AddCategory 添加分类
func AddCategory(req *val.GetCategoryReq) (resp string, err error) {
	if req.Name == "" {
		err = mistake.NewReqErr(mistake.ErrUnknown, "分类名称不能为空")
		return
	}

	category := model.Category{
		Name: req.Name,
		Icon: req.Icon,
	}

	err = dao.AddCategory(&category)
	if err != nil {
		return "", mistake.New500ServiceErr(mistake.ErrUnknown, err)
	}
	return "", nil
}
