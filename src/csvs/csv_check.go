package csvs

import (
	"fmt"
	"math/rand"
)

type DropGroup struct {
	DropId      int
	WeightAll   int
	DropConfigs []*ConfigDrop
}

var (
	ConfigDropGroupMap map[int]*DropGroup
)

func CheckLoadCsv() {
	MakeDropGroupMap()
	fmt.Println("csv配置表读取完成")
}

func MakeDropGroupMap() {
	ConfigDropGroupMap = make(map[int]*DropGroup)
	for _, v := range ConfigDropSlice {
		dropGroup, ok := ConfigDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			ConfigDropGroupMap[v.DropId] = dropGroup
		}
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//RandDropTest()
	return
}

//func RandDropTest() {
//	dropGroup := ConfigDropGroupMap[1000]
//	if dropGroup == nil {
//		return
//	}
//	for {
//		config := GetRandDrop(dropGroup)
//		fmt.Println(config.Weight)
//		break
//	}
//}

func GetRandDrop(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			return v
		}
	}
	return nil
}

func GetRandDropNew(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew(dropGroup)
		}
	}
	return nil
}
