package game

import (
	"fmt"
	"gensin-server/csvs"
	"time"
)

type ShowRole struct {
	RoleId    int
	RoleLevel int
}
type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	Level          int
	PlayerLevel    int //内部接口
	PlayerExp      int
	WorldLevel     int
	WorldLevelNow  int
	WorldLevelCool int64
	Birth          int
	ShowTeam       []*ShowRole
	HideShowTeam   int
	ShowCard       []int
	//看不见的字段
	IsProhibit int //int > bool 方便扩展
	IsGM       int
} // 对于数据库某一张表

// shift + alt + f  vscode格式化

func (self *ModPlayer) SetIcon(iconId int, player *Player) {

	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端， 操作非法
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标:", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {

	if !player.ModCard.IsHasCard(cardId) {
		// 通知客户端， 操作非法
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片:", player.ModPlayer.Card)
}

func (self *ModPlayer) SetName(name string, player *Player) {
	// 调用一个HTTP地址接口判断违禁字(不好)
	if GetManageBanWord().IsBanWord(name) {
		return
	}

	player.ModPlayer.Name = name
	fmt.Println("当前名字:", player.ModPlayer.Name)
}

func (self *ModPlayer) SetSign(sign string, player *Player) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println("当前签名:", player.ModPlayer.Sign)
}

func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp
	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		if config.PlayerExp == 0 {
			break
		}
		//是否完成任务
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
	fmt.Println("当前等级:", self.PlayerLevel, "---当前经验：", self.PlayerExp)
}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("操作失败， ---当前世界等级", self.WorldLevel)
		return
	}

	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("操作失败 ---当前世界等级", self.WorldLevel, "---真实世界等级", self.WorldLevelNow)
	}

	if time.Now().Unix() < int64(self.WorldLevelCool) {
		fmt.Println("操作失败，---冷却中")
		return
	}

	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("操作成功 ---当前世界等级", self.WorldLevel, "---真实世界等级", self.WorldLevelNow)
	return
}

func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("操作失败 ---当前世界等级", self.WorldLevel, "---真实世界等级", self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("操作失败，---冷却中")
		return
	}

	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("操作成功 ---当前世界等级", self.WorldLevel, "---真实世界等级", self.WorldLevelNow)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {
	if self.Birth > 0 {
		fmt.Println("已设置过生日！")
		return
	}

	month := birth / 100
	day := birth % 100
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "月没有", day, "日")
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println(month, "月没有", day, "日")
			return
		}
	default:
		fmt.Println("没有", month, "月！")
		return
	}

	self.Birth = birth
	fmt.Println("设置成功，生日为", month, "月", day, "日")

	if self.IsBirthDay() {
		fmt.Println("今天是你的生日，生日快乐！")
	} else {
		fmt.Println("期待你的生日到来")
	}
}

func (self *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if int(month) == self.Birth/100 && day == self.Birth%100 {
		return true
	}
	return false
}

func (self *ModPlayer) SetShowCard(showCard []int, player *Player) {

	if len(showCard) > csvs.SHOW_SIZE {
		return
	}

	cardExist := make(map[int]int)
	newList := make([]int, 0)
	for _, cardId := range showCard {
		_, ok := cardExist[cardId]
		if ok {
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			continue
		}
		newList = append(newList, cardId)
		cardExist[cardId] = 1
	}
	self.ShowCard = newList
	fmt.Println(self.ShowCard)
}

func (self *ModPlayer) SetShowTeam(showRole []int, player *Player) {
	if len(showRole) > csvs.SHOW_SIZE {
		fmt.Println("消息结构错误")
		return
	}
	roleExist := make(map[int]int)
	newList := make([]*ShowRole, 0)
	for _, roleId := range showRole {
		_, ok := roleExist[roleId]
		if ok {
			continue
		}
		if !player.ModRole.IsHasRole(roleId) {
			continue
		}
		showRole := new(ShowRole)
		showRole.RoleId = roleId
		showRole.RoleLevel = player.ModRole.GetRoleLevel(roleId)
		newList = append(newList, showRole)
		roleExist[roleId] = 1
	}
	self.ShowTeam = newList
	fmt.Println(self.ShowTeam)
}

func (self *ModPlayer) SetHideShowTeam(isHide int, player *Player) {
	if isHide != csvs.LOGIC_FALSE && isHide != csvs.LOGIC_TRUE {
		return
	}
	self.HideShowTeam = isHide

}

func (self *ModPlayer) SetProhibit(prohibit int) {
	self.IsProhibit = prohibit
}

func (self *ModPlayer) SetIsGM(isGM int) {
	self.IsGM = isGM
}

func (self *ModPlayer) IsCanEnter() bool {
	return int64(self.IsProhibit) < time.Now().Unix()
}
