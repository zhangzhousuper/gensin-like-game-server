package main

import (
	"fmt"
	"server/game"
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

	fmt.Printf("数据测试 ----start\n")

	player := game.NewTestPlayer()

	player.RecvSetIcon(1) // 胡桃
	player.RecvSetIcon(2) // 温蒂
	player.RecvSetIcon(3) //钟离

	player.RecvSetCard(11) // 胡桃
	player.RecvSetCard(22) // 温蒂
	player.RecvSetCard(33) //钟离
}
