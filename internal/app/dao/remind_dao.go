package dao

import (
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/mistake"
)

// GetRemindAll 获取所有提醒
func GetRemindAll() (remind []*model.Remind, err error) {
	err = db.Engine.Find(&remind)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}
