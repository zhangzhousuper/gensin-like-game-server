package game

import "fmt"

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	Level          int
	PlayerLevel    int //内部接口
	PlayerExp      int
	WorldLevel     int
	WorldLevelCool int
	Birth          int
	ShowTeam       []int
	ShowCard       int
	//看不见的字段
	IsProhibit int //int > bool 方便扩展
	IsGM       int
} // 对于数据库某一张表

// shift + alt + f  vscode格式化

func (self *ModPlayer) SetIcon(iconId int, player *Player) {

	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端， 操作非法
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标:", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {

	if !player.ModCard.IsHasCard(cardId) {
		// 通知客户端， 操作非法
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片:", player.ModPlayer.Card)
}

func (self *ModPlayer) SetName(name string, player *Player) {
	// 调用一个HTTP地址接口判断违禁字(不好)
	if GetManageBanWord().IsBanWord(name) {
		return
	}

	player.ModPlayer.Name = name
	fmt.Println("当前名字:", player.ModPlayer.Name)
}

func (self *ModPlayer) SetSign(sign string, player *Player) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println("当前签名:", player.ModPlayer.Sign)
}
