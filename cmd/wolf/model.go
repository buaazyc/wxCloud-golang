package wolf

// Game 每一局游戏
type Game struct {
}

type GameStatus int

const (
	UnknownGameStatus GameStatus = iota
	// 建房中
	Building
	// 黑夜
	Night
	// 警长竞选
	Election
	// 投票
	Voting
)

// Player 玩家
type Player struct {
	OpenID string
	Role   Role
}

type Role int

const (
	UnknownRole Role = iota
	// 狼人
	Wolf

	// 村民
	Villager

	// 预言家
	Seer
	// 女巫
	Witch
	// 猎人
	Hunter
	// 白痴
	Guard
)
