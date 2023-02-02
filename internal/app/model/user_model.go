package model

import "time"

// User 用户
type User struct {
	ID       int       `xorm:"not null pk autoincr comment('id') INT(11) id"`
	Phone    string    `xorm:"not null comment('手机号码') VARCHAR(11) phone"`
	Password string    `xorm:"not null comment('密码') VARCHAR(32) password"`
	UserId   string    `xorm:"not null comment('用户id') VARCHAR(64) user_id"`
	LoginAt  time.Time `xorm:"not null comment('登录时间') TIMESTAMP login_at"`
	LoginIp  string    `xorm:"comment('登录ip') VARCHAR(30) login_ip"`
}

// TableName 用户 表名
func (User) TableName() string {
	return "user"
}
