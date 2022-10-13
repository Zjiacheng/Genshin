package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type RoleInfo struct {
	RoleId   int
	GetTimes int
	//todo 等级 经验、、、
}
type ModRole struct {
	RoleInfo map[int]*RoleInfo
}

func (self *ModRole) IsHasRole(roleId int) bool {
	return true
}
func (self *ModRole) GetRoleLevel(roleId int) int {
	return 80
}
func (self *ModRole) AddItem(roleId int, num int64, player *Player) {
	config := csvs.GetRoleConfig(roleId)
	for i := 0; i < int(num); i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			self.RoleInfo[roleId] = data
			fmt.Printf("恭喜！获得新角色%v\n", config.ItemName)
		} else {
			//判断是否被转化为材料
			self.RoleInfo[roleId].GetTimes++
			if self.RoleInfo[roleId].GetTimes >= csvs.ADD_ROLE_TIME_NORMAL_MIN && self.RoleInfo[roleId].GetTimes <= csvs.ADD_ROLE_TIME_NORMAL_MAX {
				player.ModBag.AddItemToBag(config.Stuff, config.StuffNum, player)
				player.ModBag.AddItemToBag(config.StuffItem, config.StuffItemNum, player)
			} else {
				player.ModBag.AddItemToBag(config.MaxStuffItem, config.MaxStuffItemNum, player)
			}
		}
	}
	player.ModIcon.CheckGetIcon(roleId)
	player.ModCard.CheckGetCard(roleId)
}
