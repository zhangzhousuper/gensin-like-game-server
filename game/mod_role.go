package game

import (
	"fmt"
	"gensin-server/csvs"
)

type RoleInfo struct {
	RoleId   int
	GetTimes int
}

type ModRole struct {
	RoleInfo map[int]*RoleInfo
}

func (self *ModRole) IsHasRole(roleId int) bool {
	return true
}

func (self *ModRole) GetRoleLevel(roleId int) int {
	return 80
}

func (self *ModRole) AddItem(roleId int, num int64) {
	for i := 0; i < int(num); i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			self.RoleInfo[roleId] = data

		} else {
			// 判断实际获得东西
			fmt.Println("获得实际物品...")
			self.RoleInfo[roleId].GetTimes += 1
		}
	}
	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("获得角色", itemConfig.ItemName, "----", self.RoleInfo[roleId].GetTimes, "次")
	}

}
