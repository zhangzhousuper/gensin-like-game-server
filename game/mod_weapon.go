package game

import (
	"fmt"
	"gensin-server/csvs"
)

type Weapon struct {
	WeaponId    int
	KeyId       int
	Level       int
	Exp         int
	StarLevel   int
	RefineLevel int
	RoleId      int
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey     int
}

func (self *ModWeapon) AddItem(itemId int, num int64) {

	config := csvs.GetWeaponConfig(itemId)
	if config == nil {
		fmt.Println("配置不存在")
		return
	}
	if len(self.WeaponInfo)+int(num) > csvs.WEAPON_MAX_COUNT {
		fmt.Println("武器数量超过上限")
		return
	}

	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = itemId
		self.MaxKey++
		weapon.KeyId = self.MaxKey
		self.WeaponInfo[weapon.KeyId] = weapon
		fmt.Println("获得武器:", csvs.GetItemName(itemId), "----武器编号", weapon.KeyId)
	}
}

func (self *ModWeapon) WeaponUp(keyId int, player *Player) {
	weapon := self.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		return
	}
	weapon.Exp += 5000
	for {
		nextLevelConfig := csvs.GetWeaponLevelConfig(weaponConfig.Star, weapon.Level+1)
		if nextLevelConfig == nil {
			fmt.Println("返还武器经验:", weapon.Exp)
			weapon.Exp = 0
			break
		}
		if weapon.StarLevel < nextLevelConfig.NeedStarLevel {
			fmt.Println("返还武器经验:", weapon.Exp)
			weapon.Exp = 0
			break
		}
		if weapon.Exp < nextLevelConfig.NeedExp {
			break
		}
		weapon.Level++
		weapon.Exp -= nextLevelConfig.NeedExp
	}
	weapon.ShowInfo()
}

func (self *Weapon) ShowInfo() {
	fmt.Println(fmt.Sprintf("key:%d,Id:%d", self.KeyId, self.WeaponId))
	fmt.Println(fmt.Sprintf("当前等级:%d,当前经验:%d,当前突破等级:%d,当前精炼等级:%d",
		self.Level, self.Exp, self.StarLevel, self.RefineLevel))
}
