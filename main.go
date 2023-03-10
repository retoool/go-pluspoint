package main

import (
	"fmt"
	cron "github.com/robfig/cron/v3"
	"go-pluspoint/utils"
	"time"
)

func main() {
	fmt.Println(utils.KairosdbHost)
	ps, err := ReadPlusPointsConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	c := cron.New()

	//每分钟运行一次
	_, err = c.AddFunc("* * * * *", func() {
		nowTime := time.Now()
		ps.OneMinute(nowTime)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//每五分钟运行一次
	_, err = c.AddFunc("*/5 * * * *", func() {
		nowTime := time.Now()
		ps.FiveMinutes(nowTime)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//每十分钟运行一次
	_, err = c.AddFunc("*/10 * * * *", func() {
		nowTime := time.Now()
		ps.TenMinutes(nowTime)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//每小时运行一次
	_, err = c.AddFunc("0 * * * *", func() {
		nowTime := time.Now()
		ps.OneHour(nowTime)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//每天运行一次
	_, err = c.AddFunc("0 0 * * *", func() {
		nowTime := time.Now()
		ps.OneDay(nowTime)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Run()
}
