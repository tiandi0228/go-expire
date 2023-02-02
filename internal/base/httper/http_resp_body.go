package httper

// RespBody 响应结构体
type RespBody struct {
	// 状态码
	Code int `json:"code"`
	// 描述
	Message string `json:"message"`
	// 内容
	Data interface{} `json:"data"`
}

// NewRespBody 创建返回数据
func NewRespBody(code int, msg string) *RespBody {
	return &RespBody{
		Code:    code,
		Message: msg,
	}
}

// NewRespBodyData 创建返回数据
func NewRespBodyData(code int, msg string, data interface{}) *RespBody {
	return &RespBody{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
