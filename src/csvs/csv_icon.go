package csvs

import "YuanShen_Server/src/utils"

type ConfigIcon struct {
	IconId  int `json:"IconId"`
	CheckId int `json:"Check"`
}

var (
	ConfigIconMap         map[int]*ConfigIcon
	ConfigIconMapByRoleId map[int]*ConfigIcon
)

func init() {
	//定义map以后一定要初始化
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)
	ConfigIconMapByRoleId = make(map[int]*ConfigIcon)
	for _, v := range ConfigIconMap {
		ConfigIconMapByRoleId[v.CheckId] = v
	}
	return
}

// GetIconConfig 将指定iconId的对应的map数据传递出去
func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func GetIconConfigByRoleId(roleId int) *ConfigIcon {
	return ConfigIconMapByRoleId[roleId]
}
