package csvs

import "YuanShen_Server/src/utils"

type ConfigDrop struct {
	DropId int `json:"DropId"`
	Weight int `json:"Weight"`
	Result int `json:"Result"`
	IsEnd  int `json:"IsEnd"`
}

// ConfigDropSlice 这里是一个切片，不是MAP
var ConfigDropSlice []*ConfigDrop

func init() {
	ConfigDropSlice = make([]*ConfigDrop, 0)
	utils.GetCsvUtilMgr().LoadCsv("Drop", &ConfigDropSlice)
	return
}
