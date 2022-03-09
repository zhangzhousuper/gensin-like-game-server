package main

import (
	"fmt"
	"gensin-server/csvs"
	"gensin-server/game"
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

	playerGM := game.NewTestPlayer()
	playerGM.ModPlayer.SetShowTeam([]int{1001, 1001, 1001, 1001, 1001, 1002, 1001, 1005, 1001, 1001, 1001, 1002, 1001, 1005}, playerGM)

	//playerGM.ModPlayer.AddExp(10000000, playerGM)

	// ticker := time.NewTicker(time.Second * 1)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		if time.Now().Unix()%3 == 0 {
	// 			playerGM.ReduceWorldLevel()
	// 		} else if time.Now().Unix()%5 == 0 {
	// 			playerGM.ReturnWorldLevel()
	// 		}
	// 	}
	// }
	return
}
