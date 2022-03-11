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

	//背包模块
	// 1 物品识别
	// 2 物品增加
	// 3 物品消耗
	// 4 物品使用
	// 5 角色模块 -> 头像模块

	// 加载配置
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()

	fmt.Printf("数据测试 ----start\n")
	playerTest := game.NewTestPlayer()
	playerTest.ModBag.AddItem(1000003)
	playerTest.ModBag.AddItem(1000006)
	playerTest.ModBag.AddItem(1000008)
	playerTest.ModBag.AddItem(2000002)
	playerTest.ModBag.AddItem(2000021)
	playerTest.ModBag.AddItem(2000088)
	playerTest.ModBag.AddItem(3000004)
	playerTest.ModBag.AddItem(4000025)

	// ticker := time.NewTicker(time.Second * 10)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		playerTest := game.NewTestPlayer()
	// 		go playerTest.Run()
	// 	}
	// }
	return
}
