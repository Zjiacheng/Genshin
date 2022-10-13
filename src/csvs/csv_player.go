package csvs

import "YuanShen_Server/src/utils"

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	//读取csv人物等级文件
	//utils是第三方库，其要求文件必须是csv后缀且使用时不写后缀名，必须放在csv文件夹下
	utils.GetCsvUtilMgr().LoadCsv("PlayerLevel", &ConfigPlayerLevelSlice)
	return
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level > len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level-1]
}
