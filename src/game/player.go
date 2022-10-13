//  对外的接口，负责与玩家交互，如改变名字签名等
package game

const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1
	TASK_STATE_FINISH = 2
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
	ModRole       *ModRole
	ModBag        *ModBag
	ModTeam       *ModTeam
	ModWeapon     *ModWeapon
	ModRelics     *ModRelics
	ModCook       *ModCook
	ModPool       *ModPool
	ModMap        *ModMap
}

func NewTestPlayer() *Player {
	//模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	//player.ModIcon.IconInfo = make(map[int]*Icon)
	player.ModCard = new(ModCard)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModRole = new(ModRole)
	player.ModBag = new(ModBag)
	player.ModTeam = new(ModTeam)
	player.ModWeapon = new(ModWeapon)
	player.ModRelics = new(ModRelics)
	player.ModCook = new(ModCook)
	player.ModPool = new(ModPool)
	player.ModMap = new(ModMap)
	//数据初始化
	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.WorldLevel = 6
	player.ModPlayer.WorldLevelNow = 6
	player.ModIcon.IconInfo = make(map[int]*Icon)
	player.ModCard.CardInfo = make(map[int]*Card)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)
	player.ModRole.RoleInfo = make(map[int]*RoleInfo)
	player.ModWeapon.WeaponInfo = make(map[int]*Weapon)
	player.ModRelics.RelicsInfo = make(map[int]*Relics)
	player.ModCook.CookInfo = make(map[int]*Cook)
	player.ModPool.UpPoolInfo = new(PoolInfo)
	player.ModMap.InitData()
	//player.ModMap.MapInfo = make(map[int]*Map)
	return player
}

//对外接口
func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}
func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}
func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}
func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}
func (self *Player) ReturnWorldLevel() {
	self.ModPlayer.ReturnWorldLevel(self)
}
func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self) //将birth和self传递给内接口的方法
}
func (self *Player) SetShowCard(showcard []int) {
	self.ModPlayer.SetShowCard(showcard, self)
}
func (self *Player) SetShowTeam(showRole []int) {
	self.ModPlayer.SetShowTeam(showRole, self)
}
func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}
func (self *Player) Run() {
	self.ModPool.DoUpPool()
	//ticker := time.NewTicker(time.Second * 1)
	//for {
	//	select {
	//	case <-ticker.C:
	//		if time.Now().Unix()%3 == 0 {
	//			self.ModBag.UseItem(8000002, 1, self)
	//		} else {
	//			self.ModBag.AddItem(8000002, 2, self)
	//		}
	//	}
	//}
}
