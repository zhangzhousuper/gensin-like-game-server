package game

import (
	"fmt"
	"gensin-server/csvs"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (self *ModBag) AddItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println("物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.AddItemToBag(itemId, 1)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("头像", itemConfig.ItemName)
		player.ModIcon.AddItem(itemId, player)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 12)
	default: // if too many items
		//self.AddItemToBag(itemId, 1)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum += num
		fmt.Println("获得物品,数量", itemId, num)
	} else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId:  itemId,
			ItemNum: num,
		}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得物品", config.ItemName, "----数量", num)
	}
}

func (self *ModBag) RemoveItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemFromBagGM(itemId, 1)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("头像", itemConfig.ItemName)
		player.ModIcon.AddItem(itemId, player)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 12)
	default: // if too many items
		//self.AddItemToBag(itemId, 1)
	}
}

func (self *ModBag) RemoveItemFromBagGM(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	} else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId:  itemId,
			ItemNum: 0 - num,
		}
	}

	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("扣除物品", config.ItemName, "----数量", num, "当前数量", self.BagInfo[itemId].ItemNum)
	}
}
