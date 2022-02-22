package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

func NewTestPlayer() *Player {
	// 模块的初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)
	// -----------------------------
	//数据的初始化

	return player
}

// 对外接口
func (self *Player) RecvSetIcon(iconId int) {
	//Recv* 与客户端打交道的函数
	self.ModPlayer.SetIcon(iconId, self)
}

// 对外接口
func (self *Player) RecvSetCard(cardId int) {
	//Recv* 与客户端打交道的函数
	self.ModPlayer.SetCard(cardId, self)
}
