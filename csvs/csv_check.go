package csvs

import "fmt"

var (
	ConfigDropGroupMap map[int]*DropGroup
)

type DropGroup struct {
	DropId     int
	WeightAll  int
	DropConfig []*ConfigDrop
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
		dropGroup.DropConfig = append(dropGroup.DropConfig, v)
	}
	return
}
