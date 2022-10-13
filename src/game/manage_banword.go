package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
	"regexp"
	"time"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"元神", "外挂"}
	}
	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		//fmt.Println(match, v)
		if match {
			return match
		}
	}
	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		//fmt.Println(match, v)
		if match {
			return match
		}
	}
	return false
}

func (self *ManageBanWord) Run() {
	self.BanWordBase = csvs.GetBanWordBase()
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%10 == 0 {
				fmt.Println("更新词库...")
			} else {
				fmt.Println("待机...")
			}
		}
	}
}
