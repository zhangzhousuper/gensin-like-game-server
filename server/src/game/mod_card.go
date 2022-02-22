package game

type ModCard struct {
	CardId int
} // 对于数据库某一张表

func (self *ModCard) IsHasCard(iconId int) bool {
	// icon 是否存在
	return true
}
