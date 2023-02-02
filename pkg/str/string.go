package str

import "encoding/json"

// ToJSONString 转换数据为json string
func ToJSONString(data interface{}) string {
	if bytes, err := json.Marshal(data); err != nil {
		return "{}" // output empty json object string if got an error
	} else {
		return string(bytes)
	}
}
