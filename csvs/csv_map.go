package csvs

import "gensin-server/utils"

type ConfigMap struct {
	MapId   int    `json:"MapId"`
	MapName string `json:"MapName"`
}

type ConfigMapEvent struct {
	EventId     int    `json:"EventId"`
	EventType   int    `json:"EventType"`
	Name        string `json:"Name"`
	RefreshType int    `json:"RefreshType"`
	EventDrop   int    `json:"EventDrop"`
	MapId       int    `json:"MapId"`
}

var (
	ConfigMapMap      map[int]*ConfigMap
	ConfigMapEventMap map[int]*ConfigMapEvent
)

func init() {
	ConfigMapMap = make(map[int]*ConfigMap)
	utils.GetCsvUtilMgr().LoadCsv("Map", &ConfigMapMap)

	ConfigMapEventMap = make(map[int]*ConfigMapEvent)
	utils.GetCsvUtilMgr().LoadCsv("MapEvent", &ConfigMapEventMap)
	return
}

func GetMapName(mapId int) string {
	_, ok := ConfigMapMap[mapId]
	if !ok {
		return ""
	}
	return ConfigMapMap[mapId].MapName
}

func GetEventName(eventId int) string {
	_, ok := ConfigMapEventMap[eventId]
	if !ok {
		return ""
	}
	return ConfigMapEventMap[eventId].Name
}

func GetEventConfig(eventId int) *ConfigMapEvent {
	return ConfigMapEventMap[eventId]
}
