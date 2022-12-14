package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
)

type PoolInfo struct {
	PoolId        int
	FiveStarTimes int
	FourStarTimes int
	IsMustUp      int
}

type ModPool struct {
	UpPoolInfo *PoolInfo
}

func (self *ModPool) AddTimes() {
	self.UpPoolInfo.FiveStarTimes++
	self.UpPoolInfo.FourStarTimes++
}

func (self *ModPool) DoUpPool() {
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	resultEach := make(map[int]int)
	resultEachTest := make(map[int]int)
	fiveTest := 0
	for i := 0; i < 100000000; i++ {
		self.AddTimes()
		if i%10 == 0 {
			fiveTest = 0
		}
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if self.UpPoolInfo.FiveStarTimes > csvs.FIVE_STAR_TIMES_LIMIT || self.UpPoolInfo.FourStarTimes > csvs.FOUR_STAR_TIMES_LIMIT {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (self.UpPoolInfo.FiveStarTimes - csvs.FIVE_STAR_TIMES_LIMIT) * csvs.FIVE_STAR_TIMES_LIMIT_EACH_VALUE
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (self.UpPoolInfo.FourStarTimes - csvs.FOUR_STAR_TIMES_LIMIT) * csvs.FOUR_STAR_TIMES_LIMIT_EACH_VALUE
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					fiveTest++
					resultEach[self.UpPoolInfo.FiveStarTimes]++
					self.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("????????????")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
					}
				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			weaponConfig := csvs.GetWeaponConfig(roleIdConfig.Result)
			if weaponConfig != nil {
				if weaponConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
		}
		if i%10 == 9 {
			resultEachTest[fiveTest]++
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("??????%s?????????%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("??????4????????????%d", fourNum))
	fmt.Println(fmt.Sprintf("??????5??????%d", fiveNum))

	for k, v := range resultEach {
		fmt.Println(fmt.Sprintf("???%d?????????5???????????????%d", k, v))
	}

	for k, v := range resultEachTest {
		fmt.Println(fmt.Sprintf("10???%d????????????%d", k, v))
	}
}
