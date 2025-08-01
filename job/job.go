package job

import (
	"fmt"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func CreateJob(db *gorm.DB) {
	fmt.Println("创建定时任务")
	// 启动定时任务
	c := cron.New()
	// 挖矿，每天凌晨1点0分执行一次挖矿操作
	err := c.AddFunc("0 0 1 * * *", func() {
		fmt.Println("执行第挖矿定时任务")
		// 执行挖矿
	})
	if err != nil {
		return
	}

	// 开始获取价格
	//go CreatePrice(db)

	c.Start()
}
