package game

import (
	"fmt"
	"regexp"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string //从配置表做的东西
	BanWordExtra []string //等内/外部调用 来更新
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"原神"}
	}
	return manageBanWord
} // 单例好处就是任何时候返回的都是唯一的东西

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return true
		}
	}

	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return true
		}
	}
	return false
}
