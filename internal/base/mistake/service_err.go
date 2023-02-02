package mistake

import (
	"fmt"
	"hongcha/go-expire/internal/base/logger"
	"net/http"
)

// ServiceErr 业务服务异常
type ServiceErr struct {
	Code     int
	HTTPCode int
	Message  string
	Err      error
	Stack    string
}

func (e *ServiceErr) Error() string {
	return e.Message
}

// New500ServiceErr 服务器异常错误
func New500ServiceErr(code int, err error) *ServiceErr {
	return &ServiceErr{
		Code:     code,
		HTTPCode: http.StatusInternalServerError,
		Message:  GetErrInfo(code),
		Err:      err,
		Stack:    logger.LogStack(2, 5),
	}
}

// New400ServiceErr 业务错误
func New400ServiceErr(err error, message string) *ServiceErr {
	return &ServiceErr{
		HTTPCode: http.StatusBadRequest,
		Message:  message,
		Err:      err,
		Stack:    logger.LogStack(2, 5),
	}
}

// New400ServiceErrCode 构建 400 错误
func New400ServiceErrCode(code int) *ServiceErr {
	return &ServiceErr{
		Code:     code,
		HTTPCode: http.StatusBadRequest,
		Message:  GetErrInfo(code),
		Err:      fmt.Errorf(GetErrInfo(code)),
		Stack:    logger.LogStack(2, 5),
	}
}

// NewServiceErr 构建 Diy 错误
func NewServiceErr(code, httpCode int, err error, message string) *ServiceErr {
	return &ServiceErr{
		Code:     code,
		HTTPCode: httpCode,
		Message:  message,
		Err:      err,
		Stack:    logger.LogStack(2, 5),
	}
}
