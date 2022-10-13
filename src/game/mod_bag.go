package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	//case csvs.ITEMTYPE_NORMAL:
	//	self.AddItemToBag(itemId, num)
	case csvs.ITEMTYPE_ROLE:
		player.ModRole.AddItem(itemId, num, player)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 10)
	case csvs.ITEMTYPE_WEAPON:
		player.ModWeapon.AddItem(itemId, num)
	case csvs.ITEMTYPE_RELICS:
		player.ModRelics.AddItem(itemId, num)
	case csvs.ITEMTYPE_COOK:
		player.ModCook.AddItem(itemId)

	default:
		self.AddItemToBag(itemId, num, player)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64, player *Player) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum += num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Printf("获得物品----%v*%v,当前该物品总数:%v\n", config.ItemName, num, self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) RemoveItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemToBag(itemId, num)
	default:

	}
}

func (self *ModBag) RemoveItemToBag(itemId int, num int64) {
	if !self.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
			fmt.Println("物品数量不足，无法扣除!当前数量为:", nowNum)
		}
		return
	}
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: -num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Printf("减少物品%v的数量为%v,当前该物品总数:%v\n", config.ItemName, num, self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	_, ok := self.BagInfo[itemId]
	if !ok {
		return false
	} else if self.BagInfo[itemId].ItemNum < num {
		return false
	}
	return true
}

func (self *ModBag) UseItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}

	if !self.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
			fmt.Println("物品数量不足，无法扣除!当前数量为:", nowNum)
		}
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_COOKBOOK:
		self.UseCookBook(itemId, num, player)
	case csvs.ITEMTYPE_FOOD:
		fmt.Println("...")
	default:
		fmt.Println("非消耗类物品无法被使用:", csvs.GetItemName(itemId))
	}
}

func (self *ModBag) UseCookBook(itemId int, num int64, player *Player) {
	cookBookConfig := csvs.GetCookBookConfig(itemId)
	if cookBookConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	self.RemoveItem(itemId, num, player)
	self.AddItem(cookBookConfig.Reward, num, player)
}
