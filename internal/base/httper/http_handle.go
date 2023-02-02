package httper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/base/logger"
	"hongcha/go-expire/internal/base/mistake"
	validation "hongcha/go-expire/internal/base/validator"
	"hongcha/go-expire/pkg/str"
	"net/http"
)

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		c.JSON(http.StatusOK, NewRespBodyData(mistake.ErrSuccess, mistake.GetErrInfo(mistake.ErrSuccess), data))
		return
	}
	// 处理请求参数异常
	var reqErr *mistake.ReqErr
	if errors.As(err, &reqErr) {
		c.JSON(http.StatusBadRequest, NewRespBody(reqErr.Code, reqErr.Message))
		return
	}
	// 处理服务层异常
	var serviceErr *mistake.ServiceErr
	if errors.As(err, &serviceErr) {
		if serviceErr.HTTPCode >= http.StatusInternalServerError {
			logger.Error(serviceErr.Message, "\n", serviceErr.Err, "\n", serviceErr.Stack)
		} else {
			logger.Warn(serviceErr.Message, "\n", serviceErr.Err, "\n", serviceErr.Stack)
		}
		c.JSON(serviceErr.HTTPCode, NewRespBody(serviceErr.Code, serviceErr.Message))
		return
	}
	// 处理没有鉴权的异常
	var unAuthErr *mistake.StatusUnauthorizedErr
	if errors.As(err, &unAuthErr) {
		c.JSON(http.StatusUnauthorized, NewRespBody(unAuthErr.Code, unAuthErr.Message))
		return
	}
	// 处理数据库异常
	var daoErr *mistake.DaoErr
	if errors.As(err, &daoErr) {
		logger.Error(daoErr.Err, "\n", daoErr.Stack)
		c.JSON(http.StatusInternalServerError,
			NewRespBody(mistake.ErrDatabase, mistake.GetErrInfo(mistake.ErrDatabase)))
		return
	}
	// 未知错误
	logger.Error(err, logger.LogStack(2, 5))
	c.JSON(http.StatusInternalServerError, NewRespBody(mistake.ErrUnknown, mistake.GetErrInfo(mistake.ErrUnknown)))
}

// BindAndCheck 绑定参数并验证参数
// true: 确实存在问题，会返回400错误
// false: 不存在问题，验证通过
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	// 参数映射
	if err := ctx.ShouldBind(data); err != nil {
		// 当入参格式不正确的会出现，比如int传递为string
		logger.Errorf("http_handle BindAndCheck fail, %s", err.Error())
		HandleResponse(ctx, mistake.NewReqErr(mistake.ErrBind, mistake.GetErrInfo(mistake.ErrBind)), nil)
		return true
	}
	// 去除结构体内部数据前后的空格
	str.TrimStruct(data)
	// 验证参数
	if err := validation.GlobalValidate.Check(data); err != nil {
		HandleResponse(ctx, mistake.NewReqErr(mistake.ErrValidation, err.Error()), nil)
		return true
	}
	return false
}
