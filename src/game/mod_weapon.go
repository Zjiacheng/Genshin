package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type Weapon struct {
	WeaponId int
	KeyId    int
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey     int
}

func (self *ModWeapon) AddItem(itemId int, num int64) {
	if len(self.WeaponInfo)+int(num) > csvs.WEAPEN_MAX_NUM {
		fmt.Println("添加失败:超过背包容量")
		return
	}
	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = itemId
		self.MaxKey++
		weapon.KeyId = self.MaxKey
		self.WeaponInfo[weapon.KeyId] = weapon
		fmt.Printf("获得武器%v----武器编号:%v\n", csvs.GetItemName(itemId), weapon.KeyId)
	}
}
