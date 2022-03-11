package csvs

import (
	"gensin-server/utils"
)

const (
	ITEMTYPE_NORMAL = 1
	ITEMTYPE_ROLE   = 2
	ITEMTYPE_ICON   = 3
	ITEMTYPE_CARD   = 4
)

type ConfigItem struct {
	ItemId   int    `json:"ItemId"`
	SortType int    `json:"SortType"`
	ItemName string `json:"ItemName"`
}

var (
	ConfigItemMap map[int]*ConfigItem
)

func init() {
	ConfigItemMap = make(map[int]*ConfigItem)
	utils.GetCsvUtilMgr().LoadCsv("Item", &ConfigItemMap)
	return
}

func GetItemConfig(itemId int) *ConfigItem {
	return ConfigItemMap[itemId]
}
