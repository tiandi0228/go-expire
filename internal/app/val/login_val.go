package val

import "time"

// GetLoginReq 登录 请求结构体
type GetLoginReq struct {
	// 手机号码
	Phone string `json:"phone"`
	// 密码
	Password string `json:"password"`
	// ip
	Ip string `json:"ip"`
}

// GetLoginResp 登录返回结构体
type GetLoginResp struct {
	// 手机号码
	Phone string `json:"phone"`
	// 用户id
	UserId string `json:"user_id"`
	// 登录时间
	LoginAt time.Time `json:"login_at"`
	// token
	AccessToken string `json:"access_token"`
}
