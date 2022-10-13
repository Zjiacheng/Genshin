//内部接口文件，即玩家不能直接调用的接口，如获得经验等
package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
	"time"
)

//定义数据库表结构
type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelNow  int
	WorldLevelCool int64
	Birth          int
	ShowTeam       []*ModTeam
	HideShowTeam   int //是否隐藏阵容 0和1表示
	ShowCard       []int
	Prohibit       int
	IsGM           int
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if player.ModIcon.IsHasIcon(iconId) != true {
		//操作非法直接终止
		fmt.Println("不存在头像:", iconId)
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Printf("修改成功，当前图标: %d\n", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if player.ModCard.IsHasCard(cardId) != true {
		//操作非法直接终止
		fmt.Println("不存在名片:", cardId)
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Printf("修改成功，当前名片: %d\n", player.ModPlayer.Card)
}

func (self *ModPlayer) SetName(name string, player *Player) {
	if GetManageBanWord().IsBanWord(name) {
		fmt.Printf("修改失败，名称中包含违禁字符！\n")
		return
	}
	player.ModPlayer.Name = name
	fmt.Printf("修改成功，当前名称: %v\n", player.ModPlayer.Name)
}

func (self *ModPlayer) SetSign(sign string, player *Player) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Printf("修改成功，当前签名: %v\n", player.ModPlayer.Sign)
}

func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp

	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			fmt.Println("读取经验配置表失败，请仔细看看吧！")
			break
		}
		if config.PlayerExp == 0 {
			fmt.Printf("已达到最高等级咯，请去休息会儿吧！")
			break
		}
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			break
		}
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}
	fmt.Printf("当前等级: %v\n", self.PlayerLevel)
}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_STATE {
		fmt.Printf("当前世界等级为%v，不能降低\n", self.WorldLevel)
		return
	}

	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Printf("当前世界等级为%v，真实世界等级为%v,不能降低\n", self.WorldLevel, self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("世界等级降低功能还在冷却中，请间隔一天再点击哦")
	}

	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL
	fmt.Printf("降低世界等级成功，当前世界等级为%v\n", self.WorldLevelNow)
	return
}

func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("当前为最大世界等级")
	}
	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("世界等级降低功能还在冷却中，请间隔一天再点击哦")
	}
	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL
	fmt.Printf("还原世界等级成功，当前世界等级为%v\n", self.WorldLevelNow)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {
	month := birth / 100
	day := birth % 100
	//switch-case结构
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Printf("日期输入不合法，%v月没有%v天", month, day)
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Printf("日期输入不合法，%v月没有%v天", month, day)
			return
		}
	case 2:
		if day <= 0 || day > 28 {
			fmt.Printf("日期输入不合法，%v月没有%v天", month, day)
			return
		}
	default:
		fmt.Println("月份输入不合法")
	}
	self.Birth = birth
	fmt.Printf("生日设置成功，您的生日为%v月%v日\n", month, day)
}

func (self *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if int(month) == self.Birth/100 && day == self.Birth%100 {
		fmt.Println("Happy Birthday!")
		return true
	}
	return false
}

func (self *ModPlayer) SetShowCard(showcard []int, player *Player) {
	cardExist := make(map[int]int)
	newList := make([]int, 0)
	for _, cardId := range showcard { //_表示序号，cardId是我们需要的值
		_, ok := cardExist[cardId] //键、值，查看对应键的值是否存在
		if ok {
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			continue
		}
		if len(showcard) > 9 {
			//属于违法操作，直接终止
			return
		}
		newList = append(newList, cardId)
		cardExist[cardId] = 1
	}
	self.ShowCard = newList
	fmt.Println(self.ShowCard)
}

func (self *ModPlayer) SetShowTeam(showRole []int, player *Player) {
	roleExist := make(map[int]int)
	newList := make([]*ModTeam, 0)
	for _, roleId := range showRole {
		_, ok := roleExist[roleId]
		if ok {
			continue
		}
		if player.ModTeam.IsExistRole(roleId) != true {
			fmt.Println("未拥有该角色，不能设置角色名片")
			return
		}
		if len(showRole) > 6 {
			fmt.Println("请求不合法，超过展示栏位")
		}
		roleExist[roleId] = 1
		showRole := new(ModTeam)
		showRole.RoleId = roleId
		showRole.RoleLevel = player.ModTeam.GetRoleLevel(roleId)
		newList = append(newList, showRole)
	}
	self.ShowTeam = newList
	fmt.Println(self.ShowTeam[0])
}

func (self *ModPlayer) SetHideShowTeam(isHide int, player *Player) {
	if isHide != csvs.LOGIC_FALSE && isHide != csvs.LOGIC_TRUE {
		return
	}
	self.HideShowTeam = isHide
}
func (self *ModPlayer) SetProhibit(prohibit int) {
	self.Prohibit = prohibit
}
func (self *ModPlayer) SetIsGM(isGm int) {
	self.IsGM = isGm
}
func (self *ModPlayer) IsCanEnter() bool {
	return int64(self.Prohibit) < time.Now().Unix()
}
