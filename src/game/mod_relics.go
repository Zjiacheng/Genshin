package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Relics struct {
	RelicsId int
	KeyId    int
}

type ModRelics struct {
	RelicsInfo map[int]*Relics
	MaxKey     int
}

func (self *ModRelics) AddItem(itemId int, num int64) {
	if len(self.RelicsInfo)+int(num) > csvs.Relics_MAX_NUM {
		fmt.Println("添加失败:超过背包容量")
		return
	}
	for i := int64(0); i < num; i++ {
		relics := new(Relics)
		relics.RelicsId = itemId
		self.MaxKey++
		relics.KeyId = self.MaxKey
		self.RelicsInfo[relics.KeyId] = relics
		fmt.Printf("获得圣遗物%v----圣遗物编号:%v\n", csvs.GetItemName(itemId), relics.KeyId)
	}

}
