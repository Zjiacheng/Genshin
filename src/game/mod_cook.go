package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Cook struct {
	CookId int
}

type ModCook struct {
	CookInfo map[int]*Cook
}

func (self *ModCook) AddItem(itemId int) {
	_, ok := self.CookInfo[itemId]
	if ok {
		fmt.Println("已习得:", csvs.GetItemName(itemId))
		return
	}
	config := csvs.GetItemConfig(itemId)
	if config == nil {
		fmt.Println("没有该烹饪技能", csvs.GetItemName(itemId))
		return
	}
	self.CookInfo[itemId] = &Cook{itemId}
	fmt.Println("恭喜习得烹饪:", csvs.GetItemName(itemId))
}
