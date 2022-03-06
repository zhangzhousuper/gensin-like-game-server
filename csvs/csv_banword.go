package csvs

import (
	"fmt"
)

type ConfigBanWord struct {
	Id  int
	Txt string
}

var ConfigBanWordSlice []*ConfigBanWord

func init() {
	//Loadcsv(ConfigBanWord, "banword.csv") // 正常写法
	ConfigBanWordSlice = append(ConfigBanWordSlice,
		&ConfigBanWord{Id: 1, Txt: "外挂"},
		&ConfigBanWord{Id: 2, Txt: "辅助"},
		&ConfigBanWord{Id: 3, Txt: "微信"},
		&ConfigBanWord{Id: 4, Txt: "代练"},
		&ConfigBanWord{Id: 5, Txt: "赚钱"},
	)

	fmt.Println("cvf_banword init")
}

func GetBanWordBase() []string {
	relString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		relString = append(relString, v.Txt)
	}
	return relString
}
