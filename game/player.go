package game

import (
	"fmt"
	"time"
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
	ModRole       *ModRole
	ModBag        *ModBag
}

func NewTestPlayer() *Player {
	// 模块的初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)
	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	player.ModRole = new(ModRole)
	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)
	// -----------------------------
	//数据的初始化
	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.WorldLevel = 6
	player.ModPlayer.WorrldLevelNow = 6

	return player
}

// 对外接口
func (self *Player) RecvSetIcon(iconId int) {
	//Recv* 与客户端打交道的函数
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}

func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}

func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}

func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}

func (self *Player) ReturnWorldLevel() {
	self.ModPlayer.ReturnWorldLevel(self)
}

func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self)
}

func (self *Player) SetShowCard(ShowCard []int) {
	self.ModPlayer.SetShowCard(ShowCard, self)
}

func (self *Player) SetShowTeam(ShowRole []int) {
	self.ModPlayer.SetShowTeam(ShowRole, self)
}

func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}

func (self *Player) Run() {
	ticker := time.NewTicker((time.Second * 1))
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
		}
	}
}
