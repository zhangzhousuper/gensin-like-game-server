package game

import (
	"encoding/json"
	"fmt"
	"gensin-server/csvs"
	"io/ioutil"
	"os"
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

	player *Player
	path   string
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

func (self *ModWeapon) WeaponUpStar(keyId int, player *Player) {
	weapon := self.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		return
	}
	nextStarConfig := csvs.GetWeaponStarConfig(weaponConfig.Star, weapon.StarLevel+1)
	if nextStarConfig == nil {
		return
	}
	//验证物品充足并扣除
	//........
	if weapon.Level < nextStarConfig.Level {
		fmt.Println("武器等级不够，无法突破")
		return
	}
	weapon.StarLevel++
	weapon.ShowInfo()
}

func (self *ModWeapon) WeaponUpRefine(keyId int, targetKeyId int, player *Player) {
	if keyId == targetKeyId {
		fmt.Println("错误的材料")
		return
	}
	weapon := self.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponTarget := self.WeaponInfo[targetKeyId]
	if weaponTarget == nil {
		return
	}
	if weapon.WeaponId != weaponTarget.WeaponId {
		fmt.Println("错误的材料")
		return
	}
	if weapon.RefineLevel >= csvs.WEAPON_MAX_REFINE {
		fmt.Println("已经达到最大精炼等级")
		return
	}
	weapon.RefineLevel++
	delete(self.WeaponInfo, targetKeyId)
	weapon.ShowInfo()
}

func (self *ModWeapon) SaveData() {
	content, err := json.Marshal(self)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(self.path, content, os.ModePerm)
	if err != nil {
		return
	}
}

func (self *ModWeapon) LoadData(player *Player) {

	self.player = player
	self.path = self.player.localPath + "/weapon.json"

	configFile, err := ioutil.ReadFile(self.path)
	if err != nil {
		self.InitData()
		return
	}
	err = json.Unmarshal(configFile, &self)
	if err != nil {
		self.InitData()
		return
	}

	if self.WeaponInfo == nil {
		self.WeaponInfo = make(map[int]*Weapon)
	}
	return
}

func (self *ModWeapon) InitData() {
	if self.WeaponInfo == nil {
		self.WeaponInfo = make(map[int]*Weapon)
	}
}
