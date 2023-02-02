package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Str 加密
func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
