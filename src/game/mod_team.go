package game

type ModTeam struct {
	RoleId    int
	RoleLevel int
}

func (self *ModTeam) IsExistRole(roleId int) bool {
	//todo
	return true
}

func (self *ModTeam) GetRoleLevel(roleId int) int {
	return 100
}
