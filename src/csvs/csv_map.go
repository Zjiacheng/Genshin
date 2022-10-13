package csvs

import "YuanShen_Server/src/utils"

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
	ConfigMapMap          map[int]*ConfigMap
	ConfigMapMapMondStadt map[int]*ConfigMapEvent
)

func init() {
	ConfigMapMap = make(map[int]*ConfigMap)
	utils.GetCsvUtilMgr().LoadCsv("Map", &ConfigMapMap)
	ConfigMapMapMondStadt = make(map[int]*ConfigMapEvent)
	utils.GetCsvUtilMgr().LoadCsv("MapMondStadt", &ConfigMapMapMondStadt)
	return
}
