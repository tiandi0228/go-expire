package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/httper"
	"hongcha/go-expire/internal/base/mistake"
	"strings"
)

func Authorize(ctx *gin.Context) {

	token := ExtractToken(ctx)

	if len(token) == 0 {
		httper.HandleResponse(ctx, mistake.NewStatusUnauthorizedErr(mistake.ErrTokenInvalid), nil)
		ctx.Abort()
		return
	}

	resp, err := db.GetCacheString(token)
	fmt.Println("token", len(resp), token)
	if err != nil || len(resp) == 0 {
		httper.HandleResponse(ctx, mistake.NewStatusUnauthorizedErr(mistake.ErrTokenInvalid), nil)
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}

// ExtractToken 从request请求中获取token
func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader("access-token")
	if len(token) == 0 {
		token = ctx.Query("access-token")
	}
	return strings.TrimPrefix(token, "Bearer ")
}
