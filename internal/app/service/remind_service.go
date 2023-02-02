package service

import (
	"github.com/jinzhu/copier"
	"hongcha/go-expire/internal/app/dao"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/mistake"
)

// GetRemindAll 获取所有提醒
func GetRemindAll() (resp []*val.GetRemindResp, err error) {
	remindAll, err := dao.GetRemindAll()
	if err != nil {
		return nil, mistake.New500ServiceErr(mistake.ErrUnknown, err)
	}

	resp = make([]*val.GetRemindResp, 0)

	_ = copier.Copy(&resp, &remindAll)

	return
}
