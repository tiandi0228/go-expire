package model

// Remind 提醒时间
type Remind struct {
	ID    int    `xorm:"not null pk autoincr comment('id') INT(11) id"`
	Label string `xorm:"not null comment('名称') VARCHAR(11) label"`
	Value int    `xorm:"not null comment('值') INT(11) value"`
}

// TableName 提醒时间 表名
func (Remind) TableName() string {
	return "remind"
}
