package str

import (
	"hongcha/go-expire/pkg/reflecter"
	"reflect"
	"strings"
)

// TrimStruct 去除结构体的空格
func TrimStruct(obj interface{}) {
	objT, objV, err := reflecter.GetPtrTypeValue(obj)
	if err != nil {
		return
	}

	switch objT.Kind() {
	case reflect.Struct:
		doTrimFields(objT, objV)
	case reflect.Slice:
		for i := 0; i < objV.Len(); i++ {
			doTrimFields(objV.Index(i).Type(), objV.Index(i))
		}
	}
}

// doTrimFields 执行去除结构体的空格
func doTrimFields(objT reflect.Type, objV reflect.Value) {
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}
		fieldT := objT.Field(i)
		kind := fieldT.Type.Kind()
		if fieldT.Anonymous && kind == reflect.Struct {
			doTrimFields(fieldT.Type, fieldV)
			continue
		}
		if !fieldV.CanInterface() {
			continue
		}
		switch kind {
		case reflect.String:
			fieldV.SetString(strings.TrimSpace(fieldV.Interface().(string)))
		case reflect.Slice:
			if ls, ok := fieldV.Interface().([]string); ok {
				res := make([]string, len(ls))
				for i, v := range ls {
					res[i] = strings.TrimSpace(v)
				}
				fieldV.Set(reflect.ValueOf(res))
			}
		}
	}
}
