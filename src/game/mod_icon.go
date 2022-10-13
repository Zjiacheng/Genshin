package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Icon struct {
	IconId int
}

// ModIcon 验证该头像是否已经在使用
type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := self.IconInfo[iconId]
	return ok //这里的ok：IconInfo中有该头像则为True，没有则为False
}

func (self *ModIcon) AddItem(itemId int) {
	if self.IsHasIcon(itemId) {
		return
	}
	config := csvs.GetIconConfig(itemId)
	if config == nil {
		fmt.Println("配置不存在")
		return
	}
	self.IconInfo[itemId] = &Icon{IconId: itemId}
	fmt.Println("获得头像:", itemId)
}

func (self *ModIcon) CheckGetIcon(roleId int) {
	config := csvs.GetIconConfigByRoleId(roleId)
	if config == nil {
		return
	}
	self.AddItem(config.IconId)
}
