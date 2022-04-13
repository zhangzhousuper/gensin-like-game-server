package main

import (
	"gensin-server/csvs"
	"gensin-server/game"
	"math/rand"
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

	//背包模块
	// 1 物品识别
	// 2 物品增加
	// 3 物品消耗
	// 4 物品使用
	// 5 角色模块 -> 头像模块

	//掉落组模块
	// 1 保底设计
	// 2 大数据测试 (为策划服务)
	// 3 更新测试工具
	// 4 UP池子
	// 5 仓库检查 (不认为会有)

	// 地图模块
	// 1 蒙德地图
	// 2 地图上的数据结构, 1.采集物(矿物), 2. 怪物(低级怪物,北风狼王), 3. 传送点 							4. 七天神像, 神瞳 5. 宝箱
	// 3 秘境地图: 1.圣遗物秘境, 顺便为后面圣遗物模块做支持 2.封魔龙

	// 加载配置
	rand.Seed(time.Now().Unix())
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()

	//fmt.Printf("数据测试 ----start\n")
	playerTest := game.NewTestPlayer()
	go playerTest.Run()

	for {

	}
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
