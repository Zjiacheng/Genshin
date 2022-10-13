package csvs

import "YuanShen_Server/src/utils"

type CardData struct {
	CardId       int `json:"CardId"`
	Friendliness int `json:"Friendliness"`
	Check        int `json:"Check"`
}

var (
	CardDataMap       map[int]*CardData
	CardDataMapByRole map[int]*CardData
)

func init() {
	CardDataMap = make(map[int]*CardData)
	utils.GetCsvUtilMgr().LoadCsv("Card", &CardDataMap)
	CardDataMapByRole = make(map[int]*CardData)
	for _, v := range CardDataMap {
		CardDataMapByRole[v.Check] = v
	}
	return
}

func GetCardData(cardId int) *CardData {
	return CardDataMap[cardId]
}

func GetCardDataByRole(roleId int) *CardData {
	return CardDataMapByRole[roleId]
}
