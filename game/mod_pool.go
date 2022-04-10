package game

import (
	"fmt"
	"gensin-server/csvs"
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

func (self *ModPool) DoUpPool() {
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	resultEach := make(map[int]int)
	for i := 0; i < 10000000; i++ {
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
				dropGroup = newDropGroup
			}
		}

		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					resultEach[self.UpPoolInfo.FiveStarTimes]++
					// 抽中五星后
					self.UpPoolInfo.FiveStarTimes = 0
					self.UpPoolInfo.FourStarTimes++
					fiveNum++
					// 大保底
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
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
					self.UpPoolInfo.FourStarTimes++
					fourNum++
				}
			} else {
				self.UpPoolInfo.FiveStarTimes++
				self.UpPoolInfo.FourStarTimes++
			}
			result[roleIdConfig.Result]++
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))

	for k, v := range resultEach {
		fmt.Println(fmt.Sprintf("第%d抽抽中五星的概率:%d", k, v))
	}
}
