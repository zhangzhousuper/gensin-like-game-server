package main

import (
	"fmt"
	"gensin-server/csvs"
	"gensin-server/game"
	"time"
)

func main() {

	// 基础信息模块

	// 1 UID
	// 2 icon 名片
	// 3 signature
	// 4 name
	// 5 冒险等级 冒险阅历
	// 6 世界等级 冷却时间
	// 7 生日
	// 8 展示阵容 展示名片

	// 加载配置
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()

	fmt.Printf("数据测试 ----start\n")

	//playerGM.ModPlayer.AddExp(10000000, playerGM)

	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			playerTest := game.NewTestPlayer()
			go playerTest.Run()
		}
	}
	return
}
