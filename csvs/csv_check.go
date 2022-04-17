package csvs

import (
	"fmt"
	"math/rand"
)

var (
	ConfigDropGroupMap     map[int]*DropGroup
	ConfigDropItemGroupMap map[int]*DropItemGroup
	ConfigStatueMap        map[int]map[int]*ConfigStatue
)

type DropGroup struct {
	DropId      int
	WeightAll   int
	DropConfigs []*ConfigDrop
}

type DropItemGroup struct {
	DropId      int
	DropConfigs []*ConfigDropItem
}

func CheckLoadCsv() {
	// double check
	MakeDropGroupMap()
	MakeDropItemGroupMap()
	MakeConfigStatueMap()
	fmt.Println("csv init finished")
}

func MakeDropGroupMap() {
	ConfigDropGroupMap = make(map[int]*DropGroup)
	for _, v := range ConfigDropSlice {
		dropGroup, ok := ConfigDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			ConfigDropGroupMap[v.DropId] = dropGroup
		}
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//RandDropTest()
	return
}

func MakeDropItemGroupMap() {
	ConfigDropItemGroupMap = make(map[int]*DropItemGroup)
	for _, v := range ConfigDropItemSlice {
		dropGroup, ok := ConfigDropItemGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropItemGroup)
			dropGroup.DropId = v.DropId
			ConfigDropItemGroupMap[v.DropId] = dropGroup
		}
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//RandDropItemTest()
	return
}

func MakeConfigStatueMap() {
	ConfigStatueMap = make(map[int]map[int]*ConfigStatue)
	for _, v := range ConfigStatueSlice {
		statueMap, ok := ConfigStatueMap[v.StatueId]
		if !ok {
			statueMap = make(map[int]*ConfigStatue)
			ConfigStatueMap[v.StatueId] = statueMap
		}
		statueMap[v.Level] = v
	}
	return
}

func RandDropItemTest() {
	dropGroup := ConfigDropItemGroupMap[2]
	if dropGroup == nil {
		return
	}
	for _, v := range dropGroup.DropConfigs {
		randNum := rand.Intn(PERCENT_ALL)
		if randNum < v.Weight {
			fmt.Println(v.ItemId)
		}
	}
	return
}

func RandDropTest() {
	dropGroup := ConfigDropGroupMap[1000]
	if dropGroup == nil {
		return
	}
	num := 0
	for {
		config := GetRandDropNew(dropGroup)
		if config.IsEnd == LOGIC_TRUE {
			fmt.Println(GetItemName(config.Result))
			num++
			dropGroup = ConfigDropGroupMap[1000]
			if num >= 100 {
				break
			} else {
				continue
			}
		}
		dropGroup = ConfigDropGroupMap[config.Result]
		if dropGroup == nil {
			break
		}
	}
}

func GetRandDrop(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			return v
		}
	}
	return nil
}

func GetRandDropNew(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew(dropGroup)
		}
	}
	return nil
}

func GetRandDropNew1(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	// check if has
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					continue
				}

				if maxGetTime < fiveInfo[config.Result] {
					maxGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					continue
				}

				if maxGetTime < fourInfo[config.Result] {
					maxGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew1(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}

func GetRandDropNew2(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fiveInfo[config.Result] {
					minGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fourInfo[config.Result] {
					minGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew2(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}

//获得配置的接口
func GetDropItemGroup(dropId int) *DropItemGroup {
	return ConfigDropItemGroupMap[dropId]
}

func GetDropItemGroupNew(dropId int) []*ConfigDropItem {
	rel := make([]*ConfigDropItem, 0)
	config := GetDropItemGroup(dropId)
	for _, v := range config.DropConfigs {
		if v.DropType == DROP_ITEM_TYPE_ITEM {
			rel = append(rel, v)
		} else if v.DropType == DROP_ITEM_TYPE_GROUP {
			configs := GetDropItemGroupNew(v.ItemId)
			rel = append(rel, configs...)
		}
	}
	return rel
}

func GetStatueConfig(statueId int, level int) *ConfigStatue {
	_, ok := ConfigStatueMap[statueId]
	if !ok {
		return nil
	}

	_, ok = ConfigStatueMap[statueId][level]
	if !ok {
		return nil
	}
	return ConfigStatueMap[statueId][level]
}
