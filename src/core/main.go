package main

import (
	"fmt"
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

	fmt.Printf("数据测试 ----start\n")

	player := game.NewTestPlayer()

	player.RecvSetName("好人")
	player.RecvSetName("坏人")
	player.RecvSetName("求外挂")
	player.RecvSetName("感觉不如原神画质")

}
