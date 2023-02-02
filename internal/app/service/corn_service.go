package service

import (
	"fmt"
	"github.com/robfig/cron"
	"hongcha/go-expire/internal/app/dao"
	"hongcha/go-expire/internal/base/logger"
	"hongcha/go-expire/internal/base/utils"
	"strconv"
	"time"
)

func RunCorn() {
	logger.Info("run corn")

	c := cron.New()

	err := c.AddFunc("0 0 12 * * ?", JudgeExpired)

	if err != nil {
		logger.Errorf("错误: %s", err)
	}

	c.Start()
}

// JudgeExpired 判断是否过期
func JudgeExpired() {
	all, err := dao.GetProductAll()
	if err != nil {
		logger.Errorf("获取所有商品错误: %s", err)
		return
	}

	// 当前时间
	nowTime := time.Now()

	for _, product := range all {
		local, _ := time.LoadLocation("Local") //设置时区
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", product.QualityGuaranteePeriod.Format("2006-01-02 15:04:05"), local)
		day := utils.SubDays(t, nowTime)

		// 过期5天删除
		if day <= 5 {
			logger.Debugf("%s, 已过期 %s 天了, 即将删除~", product.Name, day)
			err := dao.DeleteProduct(product.ID)
			if err != nil {
				return
			}
		} else if day <= 0 { // 过期设置过期状态
			logger.Debugf("%s, 已过期 %s 天了~", product.Name, day)
			err := dao.UpdateProductStatus(product.ID, 2)
			if err != nil {
				return
			}
		} else if utils.SubDays(t, nowTime) <= product.Remind { // 过期前提醒
			logger.Debugf("%s, 还有 %s 天过期~", product.Name, strconv.Itoa(day))
			err := utils.SendMessage(fmt.Sprintf("%s, 还有 %s 天过期~", product.Name, strconv.Itoa(day)))
			if err != nil {
				logger.Errorf("发送消息错误: %s", err)
				return
			}
		}
	}
}
