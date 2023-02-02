package model

// Category 分类
type Category struct {
	ID   int    `xorm:"not null pk autoincr comment('id') INT(11) id"`
	Name string `xorm:"not null comment('名称') VARCHAR(30) name"`
	Icon string `xorm:"not null comment('图标') VARCHAR(255) icon"`
}

// TableName 分类 表名
func (Category) TableName() string {
	return "category"
}
