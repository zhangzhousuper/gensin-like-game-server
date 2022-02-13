package game

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	Level          int
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelCool int
	Birth          int
	ShowTeam       []int
	ShowCard       int
	//看不见的字段
	IsProhibit int //int > bool 方便扩展
	IsGM       int
} // 对于数据库某一张表

// shift + alt + f  vscode格式化
