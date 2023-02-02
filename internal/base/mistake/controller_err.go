package mistake

// ReqErr 请求参数异常错误
type ReqErr struct {
	Code    int
	Message string
}

func (e *ReqErr) Error() string {
	return e.Message
}

// NewReqErr 创建请求参数异常错误
func NewReqErr(code int, message string) *ReqErr {
	return &ReqErr{Code: code, Message: message}
}

// NewReqErrCode 创建请求参数异常错误
func NewReqErrCode(code int) *ReqErr {
	return &ReqErr{Code: code, Message: GetErrInfo(code)}
}

// StatusUnauthorizedErr 用户鉴权失败，没有token或token过期 错误
type StatusUnauthorizedErr struct {
	Code    int
	Message string
}

func (e *StatusUnauthorizedErr) Error() string {
	return e.Message
}

// NewStatusUnauthorizedErr 创建用户鉴权失败错误
func NewStatusUnauthorizedErr(code int) *StatusUnauthorizedErr {
	return &StatusUnauthorizedErr{
		Code:    code,
		Message: GetErrInfo(code),
	}
}
