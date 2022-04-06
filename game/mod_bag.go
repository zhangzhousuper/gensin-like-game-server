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

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println("物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.AddItemToBag(itemId, num)
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
	} else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId:  itemId,
			ItemNum: num,
		}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得物品", config.ItemName, "----数量", num, "当前数量", self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) RemoveItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemFromBagGM(itemId, num)
	default: // if too many items
		//self.AddItemToBag(itemId, 1)
	}
}

func (self *ModBag) RemoveItemFromBag(itemId int, num int64) {

	if !self.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
			fmt.Println(config.ItemName, "数量不足 ---当前数量", nowNum, "需要数量", num)
		}
		return
	}

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

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	_, ok := self.BagInfo[itemId]
	if !ok {
		return false
	} else if self.BagInfo[itemId].ItemNum < num {
		return true
	}

	return true
}
