package csvs

import "YuanShen_Server/src/utils"

type ConfigUniqueTask struct {
	TaskId    int `json:"TaskId"`
	SortType  int `json:"SortType"`
	OpenLevel int `json:"OpenLevel"`
	TaskType  int `json:"TaskType"`
	Condition int `json:"Condition"`
}

var (
	ConfigUniqueTaskMap map[int]*ConfigUniqueTask
)

func init() {
	//读取csv人物等级突破条件
	//utils是第三方库，其要求文件必须是csv后缀且使用时不写后缀名，必须放在csv文件夹下
	ConfigUniqueTaskMap = make(map[int]*ConfigUniqueTask)
	utils.GetCsvUtilMgr().LoadCsv("UniqueTask", &ConfigUniqueTaskMap)

	return
}
