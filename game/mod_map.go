package game

import (
	"fmt"
	"gensin-server/csvs"
	"math/rand"
	"time"
)

type Map struct {
	MapId     int
	EventInfo map[int]*Event
}

type Event struct {
	EventId       int
	State         int
	NextResetTime int64
}

type ModMap struct {
	MapInfo map[int]*Map
}

func (self *ModMap) InitData() {
	self.MapInfo = make(map[int]*Map)

	for _, v := range csvs.ConfigMapMap {
		_, ok := self.MapInfo[v.MapId]
		if !ok {
			self.MapInfo[v.MapId] = self.NewMapInfo(v.MapId)
		}
	}

	for _, v := range csvs.ConfigMapEventMap {
		_, ok := self.MapInfo[v.MapId]
		if !ok {
			continue
		}
		_, ok = self.MapInfo[v.MapId].EventInfo[v.EventId]
		if !ok {
			self.MapInfo[v.MapId].EventInfo[v.EventId] = new(Event)
			self.MapInfo[v.MapId].EventInfo[v.EventId].EventId = v.EventId
			self.MapInfo[v.MapId].EventInfo[v.EventId].State = csvs.EVENT_START
		}
	}
}

func (self *ModMap) NewMapInfo(mapId int) *Map {
	mapInfo := new(Map)
	mapInfo.MapId = mapId
	mapInfo.EventInfo = make(map[int]*Event)
	return mapInfo
}

func (self *ModMap) GetEventList(config *csvs.ConfigMap) {
	_, ok := self.MapInfo[config.MapId]
	if !ok {
		return
	}
	for _, v := range self.MapInfo[config.MapId].EventInfo {
		self.CheckRefresh(v)
		lastTime := v.NextResetTime - time.Now().Unix()
		noticeTime := ""
		if lastTime <= 0 {
			noticeTime = "已刷新"
		} else {
			noticeTime = fmt.Sprintf("%d秒后刷新", lastTime)
		}
		fmt.Println(fmt.Sprintf("事件Id:%d,名字:%s,状态:%d,%s", v.EventId, csvs.GetEventName(v.EventId), v.State, noticeTime))
	}
}

func (self *ModMap) SetEventState(mapId int, eventId int, state int, player *Player) {
	_, ok := self.MapInfo[mapId]
	if !ok {
		fmt.Println("地图不存在")
		return
	}
	_, ok = self.MapInfo[mapId].EventInfo[eventId]
	if !ok {
		fmt.Println("事件不存在")
		return
	}
	if self.MapInfo[mapId].EventInfo[eventId].State >= state {
		fmt.Println("状态异常")
		return
	}
	eventConfig := csvs.GetEventConfig(self.MapInfo[mapId].EventInfo[eventId].EventId)
	if eventConfig == nil {
		return
	}

	self.MapInfo[mapId].EventInfo[eventId].State = state
	if state == csvs.EVENT_FINISH {
		fmt.Println("事件完成")
	}
	if state == csvs.EVENT_END {
		config := csvs.GetDropItemGroupNew(eventConfig.EventDrop)
		for _, v := range config {
			randNum := rand.Intn(csvs.PERCENT_ALL)
			if randNum < v.Weight {
				randAll := v.ItemNumMax - v.ItemNumMin + 1
				itemNum := rand.Intn(randAll) + v.ItemNumMin
				worldLevel := player.ModPlayer.GetWorldLevelNow()
				if worldLevel > 0 {
					itemNum = itemNum * (csvs.PERCENT_ALL + worldLevel*v.WorldAdd) / csvs.PERCENT_ALL
				}
				player.ModBag.AddItem(v.ItemId, int64(itemNum), player)
			}
		}
		fmt.Println("事件领取")
	}
	if state > 0 {
		switch eventConfig.RefreshType {
		case csvs.MAP_REFRESH_SELF:
			self.MapInfo[mapId].EventInfo[eventId].NextResetTime = time.Now().Unix() + csvs.MAP_REFRESH_SELF_TIME
		}
	}
}

func (self *ModMap) RefreshDay() {
	for _, v := range self.MapInfo {
		for _, v := range self.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MAP_REFRESH_DAY {
				continue
			}
			v.State = csvs.EVENT_START
		}
	}
}

func (self *ModMap) RefreshWeek() {
	for _, v := range self.MapInfo {
		for _, v := range self.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MAP_REFRESH_WEEK {
				continue
			}
			v.State = csvs.EVENT_START
		}
	}
}

func (self *ModMap) RefreshSelf() {
	for _, v := range self.MapInfo {
		for _, v := range self.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MAP_REFRESH_SELF {
				continue
			}
			if time.Now().Unix() <= v.NextResetTime {
				v.State = csvs.EVENT_START
			}
		}
	}
}

func (self *ModMap) CheckRefresh(event *Event) {
	if event.NextResetTime > time.Now().Unix() {
		return
	}
	eventConfig := csvs.GetEventConfig(event.EventId)
	if eventConfig == nil {
		return
	}

	switch eventConfig.RefreshType {
	case csvs.MAP_REFRESH_DAY:
		count := time.Now().Unix() / csvs.MAP_REFRESH_DAY_TIME
		count++
		event.NextResetTime = count * csvs.MAP_REFRESH_DAY_TIME
	case csvs.MAP_REFRESH_WEEK:
		count := time.Now().Unix() / csvs.MAP_REFRESH_WEEK_TIME
		count++
		event.NextResetTime = count * csvs.MAP_REFRESH_WEEK_TIME
	case csvs.MAP_REFRESH_SELF:
	}
}
