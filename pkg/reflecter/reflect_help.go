package reflecter

import (
	"fmt"
	"reflect"
)

// GetPtrTypeValue 反射获取 指针类型的type 和 value
// 如果不为指针或者不为结构体就会异常
func GetPtrTypeValue(obj interface{}) (objT reflect.Type, objV reflect.Value, err error) {
	objT = reflect.TypeOf(obj)
	objV = reflect.ValueOf(obj)
	if !isStructPtr(objT) && !isSlicePtr(objT) {
		return nil, reflect.Value{}, fmt.Errorf("%v must be a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()
	return objT, objV, nil
}

// isStructPtr 判断是结构体指针
func isStructPtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

// isSlicePtr 判断是切片指针
func isSlicePtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Slice
}
