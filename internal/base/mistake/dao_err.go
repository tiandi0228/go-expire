package mistake

import (
	"hongcha/go-expire/internal/base/logger"
)

// DaoErr 数据层通用异常错误
type DaoErr struct {
	Err   error
	Stack string
}

func (e *DaoErr) Error() string {
	return e.Err.Error()
}

// NewDaoErr 创建数据层通用异常错误
func NewDaoErr(err error) *DaoErr {
	return &DaoErr{
		Err:   err,
		Stack: logger.LogStack(2, 5),
	}
}
