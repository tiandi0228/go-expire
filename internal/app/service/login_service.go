package service

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/copier"
	"hongcha/go-expire/internal/app/dao"
	"hongcha/go-expire/internal/app/val"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/logger"
	"hongcha/go-expire/internal/base/mistake"
	"hongcha/go-expire/internal/base/utils"
	"net/http"
	"strconv"
	"time"
)

// Login 登录
func Login(req *val.GetLoginReq) (resp *val.GetLoginResp, err error) {

	user, err := dao.GetUser(req.Phone, utils.MD5Str(req.Password))
	if err != nil {
		return nil, mistake.NewServiceErr(mistake.ErrUnknown, http.StatusOK, err, "未知用户")
	}

	str := strconv.Itoa(int(time.Now().Unix())) + "syc.im"
	data := []byte(str)
	token := fmt.Sprintf("%x", md5.Sum(data))
	tokenStr := base64.StdEncoding.EncodeToString([]byte(token))

	// 把当前用户的token作为key存入redis 7 天
	CacheErr := db.SetCacheString(tokenStr, "expire_"+user.UserId, 168*time.Hour)
	if CacheErr != nil {
		logger.Errorf("缓存redis失败：", CacheErr)
	}
	loginAt, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	_ = dao.UpdateUser(user.ID, loginAt, req.Ip)

	resp = &val.GetLoginResp{}

	resp.AccessToken = tokenStr

	_ = copier.Copy(&resp, &user)

	return resp, nil
}
