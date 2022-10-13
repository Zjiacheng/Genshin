package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Card struct {
	CardId     int
	Friendness int
}

// ModCard 用于存放已拥有的Card
type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) IsHasCard(cardId int) bool {
	_, ok := self.CardInfo[cardId]
	return ok
}

func (self *ModCard) AddItem(cardId int, friendness int) {
	if self.IsHasCard(cardId) {
		fmt.Println("已拥有名片:", cardId)
		return
	}
	config := csvs.GetCardData(cardId)
	if config == nil {
		fmt.Println("非法名片:", cardId)
		return
	}
	if friendness < config.Friendliness {
		fmt.Println("好感度不足，快去和角色互动吧~！")
		return
	}
	self.CardInfo[cardId] = &Card{CardId: cardId}
	fmt.Println("获得名片:", cardId)
}

func (self *ModCard) CheckGetCard(roleId int) {
	config := csvs.GetCardDataByRole(roleId)
	if config == nil {
		return
	}
	self.AddItem(config.CardId, 10)
}
