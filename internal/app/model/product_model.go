package model

import "time"

// Product 商品
type Product struct {
	ID                     int       `xorm:"not null pk autoincr comment('id') INT(11) id"`
	Name                   string    `xorm:"not null comment('名称') VARCHAR(30) name"`
	ManufactureDate        time.Time `xorm:"not null comment('生产日期') TIMESTAMP manufacture_date"`
	QualityGuaranteePeriod time.Time `xorm:"not null comment('保质期') TIMESTAMP quality_guarantee_period"`
	State                  int       `xorm:"not null comment('是否过期 1：未过期 2：过期') INT(11) state"`
	CategoryID             int       `xorm:"not null comment('分类id') INT(11) category_id"`
	Remind                 int       `xorm:"not null comment('提醒时间') VARCHAR(11) remind"`
}

// TableName 商品 表名
func (Product) TableName() string {
	return "product"
}
