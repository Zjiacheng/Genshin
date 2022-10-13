// Package csvs 配置表_违禁词库
package csvs

import "fmt"

type ConfigBanWord struct {
	Id  int
	Txt string
}

var ConfigBanWordSlice []*ConfigBanWord

func init() {
	ConfigBanWordSlice = append(ConfigBanWordSlice,
		&ConfigBanWord{Id: 1, Txt: "外挂"},
		&ConfigBanWord{Id: 1, Txt: "辅助"},
		&ConfigBanWord{Id: 1, Txt: "微信"},
		&ConfigBanWord{Id: 1, Txt: "代练"},
		&ConfigBanWord{Id: 1, Txt: "赚钱"})
	fmt.Println("csv_banword初始化")
}

func GetBanWordBase() []string {
	relString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		relString = append(relString, v.Txt)
	}
	return relString
}
