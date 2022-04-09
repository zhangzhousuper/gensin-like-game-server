package game

import (
	"fmt"
	"gensin-server/csvs"
)

type Relics struct {
	RelicsId int
	KeyId    int
}

type ModRelics struct {
	RelicsInfo map[int]*Relics
	MaxKey     int
}

func (self *ModRelics) AddItem(itemId int, num int64) {

	config := csvs.GetRelicsConfig(itemId)
	if config == nil {
		fmt.Println("配置不存在")
		return
	}

	if len(self.RelicsInfo)+int(num) > csvs.Relics_MAX_COUNT {
		fmt.Println("超过最大值")
		return
	}

	for i := int64(0); i < num; i++ {
		relics := new(Relics)
		relics.RelicsId = itemId
		self.MaxKey++
		relics.KeyId = self.MaxKey
		self.RelicsInfo[relics.KeyId] = relics
		fmt.Println("获得圣遗物:", csvs.GetItemName(itemId), "------编号:", relics.KeyId)
	}
}
