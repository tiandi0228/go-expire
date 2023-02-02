package db

import (
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// Engine orm全局引擎
var Engine *xorm.Engine

// InitDB 初始化全局数据库实例
// conn: 数据库连接
func InitDB(conn string) {
	// 创建数据库连接
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}

	// 测试数据库连接
	if err = engine.Ping(); err != nil {
		panic(err)
	}

	// 数据库连接池配置
	//engine.SetConnMaxLifetime()
	//engine.SetMaxIdleConns()
	//engine.SetMaxOpenConns()

	mapper := names.GonicMapper{}
	engine.SetTableMapper(mapper)
	engine.SetColumnMapper(mapper)
	Engine = engine
}
