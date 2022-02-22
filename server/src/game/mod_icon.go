package game

type ModIcon struct {
	IconId int
} // 对于数据库某一张表

func (self *ModIcon) IsHasIcon(iconId int) bool {
	// icon 是否存在
	return true
}
