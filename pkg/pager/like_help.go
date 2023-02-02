package pager

import (
	"fmt"
	"reflect"
	"strings"
	"xorm.io/xorm"
)

// LikeHelpWithField 将条件中的所有字符串构建like查询的session （包含FieldMap中的字段）
// Condition 为查询条件，必须为结构体指针
// FieldMap 字段映射表 （需要传递数据库中的名称）
func LikeHelpWithField(session *xorm.Session, condition interface{}, fieldMap ...string) {
	likeHelp(session, condition, true, fieldMap...)
}

// LikeHelpWithoutField 将条件中的所有字符串构建like查询的session （不包含FieldMap中的字段）
// Condition 为查询条件，必须为结构体指针
// FieldMap 字段映射表 （需要传递数据库中的名称）
func LikeHelpWithoutField(session *xorm.Session, condition interface{}, fieldMap ...string) {
	likeHelp(session, condition, false, fieldMap...)
}

// likeHelp 将条件中的所有字符串构建like查询的session
// Condition 为查询条件，必须为结构体指针
// ExistInMapNeedLike true：存在于字段映射表中'会like'；false：存在于字段映射表中的字段'不会like'
// FieldMap 字段映射表
func likeHelp(session *xorm.Session, condition interface{}, existInMapNeedLike bool, fieldMap ...string) {
	if condition == nil || session == nil {
		return
	}

	// 确保传入条件是个指针
	conValue := reflect.ValueOf(condition)
	if conValue.Kind() != reflect.Ptr {
		return
	}
	conValue = conValue.Elem()

	// 如果不是结构体就返回
	if conValue.Kind() != reflect.Struct {
		return
	}
	conType := conValue.Type()

	fieldMapping := make(map[string]bool, len(fieldMap))
	for _, value := range fieldMap {
		fieldMapping[value] = true
	}

	// 查看查询条件是否有 TableName 方法，从而获取到表名
	tableName := ""
	method := conValue.MethodByName("TableName")
	if method.IsValid() {
		values := method.Call([]reflect.Value{})
		if len(values) != 0 && values[0].Type() == reflect.TypeOf("") {
			tableName = values[0].Interface().(string)
		}
	}

	// 循环处理结构体中的每个字段
	numField := conValue.NumField()
	for i := 0; i < numField; i++ {
		field := conValue.Field(i)
		filedType := conType.Field(i)

		// 如果字段类型不是字符串就返回
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		// 如果字段是空字符串也返回
		fieldValue := field.Interface().(string)
		if len(fieldValue) == 0 {
			continue
		}

		// 获取xorm的标签，从而获取到字段在数据库中的名称
		tag := filedType.Tag
		xormTag := tag.Get("xorm")
		if len(xormTag) == 0 {
			continue
		}

		// 拿到最后一个字段名称，最后一个字段默认是表名
		xormTagArr := strings.Split(xormTag, " ")
		column := xormTagArr[len(xormTagArr)-1]

		// 如果没有获取到数据库中的名称就直接继续，没办法like查询
		if len(column) == 0 {
			continue
		}

		// existInMapNeedLike true：存在于字段映射表中'会like'；false：存在于字段映射表中的字段'不会like'
		ok := fieldMapping[column]
		if existInMapNeedLike == !ok {
			continue
		}

		// 如果能获取到表名，那么给字段查询添加表名信息，保证查询字段唯一性，避免多表关联查询时字段名称重复导致ambiguous错误
		if len(tableName) != 0 {
			column = fmt.Sprintf("`%s`.`%s`", tableName, column)
		} else {
			column = fmt.Sprintf("`%s`", column)
		}

		// 构造like语句，当前默认支持全模糊，后面考虑支持指定扩展
		session.Where(column+" LIKE ?", "%"+fieldValue+"%")

		// 需要将原来查询条件中的值赋值为空字符串，避免作为查询条件时被精确查询
		field.SetString("")
	}
}
