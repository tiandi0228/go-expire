package utils

import (
	"github.com/go-resty/resty/v2"
	"hongcha/go-expire/internal/base/conf"
	"hongcha/go-expire/internal/base/logger"
)

func SendMessage(message string) (err error) {
	logger.Debugf("SendMessage: %s", message)
	r := resty.New().R()
	_, err = r.Get(conf.GlobalConfig.Message.Connection + message + "?icon=" + conf.GlobalConfig.Message.Icon + "&title=" + conf.GlobalConfig.Message.Title)
	return
}
