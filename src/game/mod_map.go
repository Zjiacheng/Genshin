package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Map struct {
	MapId     int
	EventInfo map[int]*Event
}

type Event struct {
	EventId int
	State   int
}

type ModMap struct {
	MapInfo map[int]*Map
}

func (self *ModMap) InitData() {
	self.MapInfo = make(map[int]*Map)
	for _, v := range csvs.ConfigMapMap {
		_, ok := self.MapInfo[v.MapId]
		if !ok {
			//self.MapInfo[v.MapId] = self.NewMapInfo(self.MapInfo)
		}
	}
	_, ok := self.MapInfo[1]
	if !ok {
		self.MapInfo[1] = new(Map)
		self.MapInfo[1].EventInfo = make(map[int]*Event)
	}
	for _, v := range csvs.ConfigMapMapMondStadt {
		_, ok := self.MapInfo[1].EventInfo[v.EventId]
		if !ok {
			self.MapInfo[1].EventInfo[v.EventId] = new(Event)
			self.MapInfo[1].EventInfo[v.EventId].EventId = v.EventId
			self.MapInfo[1].EventInfo[v.EventId].State = csvs.LOGIC_FALSE
		}
	}
	for _, v := range self.MapInfo[1].EventInfo {
		fmt.Println("事件:", v.EventId)
	}
}

func (self *ModMap) NewMapInfo(mapId int) *Map {
	mapInfo := new(Map)
	mapInfo.MapId = mapId
	mapInfo.EventInfo = make(map[int]*Event)
	return mapInfo
}
