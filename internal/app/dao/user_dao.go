package dao

import (
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/mistake"
	"time"
)

// GetUser 查询用户信息
func GetUser(phone string, password string) (user *model.User, err error) {
	user = &model.User{}
	_, err = db.Engine.Table("user").Where("phone = ?", phone).And("password = ?", password).Get(user)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// UpdateUser 更新用户信息
func UpdateUser(id int, loginAt time.Time, loginIp string) (err error) {
	_, err = db.Engine.ID(id).Cols("login_at, login_ip").Update(&model.User{LoginAt: loginAt, LoginIp: loginIp})
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}
