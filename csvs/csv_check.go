package csvs

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	ConfigDropGroupMap map[int]*DropGroup
)

type DropGroup struct {
	DropId      int
	WeightAll   int
	DropConfigs []*ConfigDrop
}

func CheckLoadCsv() {
	// double check
	MakeDropGroupMap()
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
	RunDropTest()
	return
}

func RunDropTest() {
	dropGroup := ConfigDropGroupMap[1000]
	if dropGroup == nil {
		return
	}

	num := 0
	for {
		config := GetRandDrop(dropGroup)
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
	rand.Seed(time.Now().Unix())
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
