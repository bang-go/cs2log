package evt

import "time"

// Meta 基础
type Meta struct {
	Time time.Time `json:"time"`
	Type string    `json:"type"`
}

// GetType is the getter fo Meta.Type
func (m Meta) GetType() string {
	return m.Type
}

// GetTime is the getter for Meta.Time
func (m Meta) GetTime() time.Time {
	return m.Time
}

func NewMeta(ti time.Time, ty string) Meta {
	return Meta{
		Time: ti,
		Type: ty,
	}
}

// Player 选手基础基础信息
type Player struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	SteamID string `json:"steam_id"`
	Side    string `json:"side"`
}

// Position 地图坐标
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// PositionFloat 精确的地图坐标
type PositionFloat struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Velocity 速率信息
type Velocity struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Equation 金钱变化的方程 A + B=Result
type Equation struct {
	A      int `json:"a"`
	B      int `json:"b"`
	Result int `json:"result"`
}

// Unknown 未定义的类型
type Unknown struct {
	Meta
	Raw string `json:"raw"`
}

// LoadingMap 加载地图
type LoadingMap struct {
	Meta
	Map string `json:"map"`
}

// ServerMessage 服务器消息
type ServerMessage struct {
	Meta
	Text string `json:"text"`
}

// Rcon Rcon服务器消息
type Rcon struct {
	Meta
	IP      string `json:"ip"`
	Port    uint   `json:"port"`
	Command string `json:"command"`
}

// FreezeTimeStart 每回合开始的冻结时间
type FreezeTimeStart struct {
	Meta
}

// WorldMatchStart 比赛开始(地图)
type WorldMatchStart struct {
	Meta
	Map string `json:"map"`
}

// WorldReloadedStart 比赛回档(地图)
type WorldReloadedStart struct {
	Meta
	Map string `json:"map"`
}

// WorldWarmupStart 热身开始
type WorldWarmupStart struct{ Meta }

// WorldWarmupEnd 热身结束
type WorldWarmupEnd struct{ Meta }

// WorldRoundStart 回合开始
type WorldRoundStart struct{ Meta }

// WorldRoundRestart 回合重新开始
type WorldRoundRestart struct {
	Meta
	TimeLeft int `json:"timeleft"`
}

// WorldKnifeStart 刀局开始
type WorldKnifeStart struct {
	Meta
}

// WorldKnifeWinner 刀局胜利方
type WorldKnifeWinner struct {
	Meta
	Side string `json:"side"`
}

// WorldKnifeChoice 刀局胜利方选择
type WorldKnifeChoice struct {
	Meta
	Side string `json:"side"`
}

// WorldRoundEnd 回合结束
type WorldRoundEnd struct{ Meta }

// WorldGameCommencing 游戏开始
type WorldGameCommencing struct{ Meta }

// WorldGameOver 游戏结束
type WorldGameOver struct {
	Meta
	Mode     string `json:"mode"`
	MapGroup string `json:"map_group"`
	Map      string `json:"map"`
	ScoreCT  int    `json:"score_ct"`
	ScoreT   int    `json:"score_t"`
	Duration int    `json:"duration"`
}

// TeamScored 回合结束后的队伍比赛
type TeamScored struct {
	Meta
	Side       string `json:"side"`
	Score      int    `json:"score"`
	NumPlayers int    `json:"num_players"`
}

// TeamNotice 回合结束的播报(队伍赢得回合)
type TeamNotice struct {
	Meta
	Side    string `json:"side"`
	Notice  string `json:"notice"`
	ScoreCT int    `json:"score_ct"`
	ScoreT  int    `json:"score_t"`
}

// MatchTeamScore 比赛队伍比分
type MatchTeamScore struct {
	Meta
	ScoreCT      int    `json:"scoreCT"`
	ScoreT       int    `json:"scoreT"`
	Map          string `json:"map"`
	RoundsPlayed int    `json:"rounds_played"`
}

// MatchTeamSide 比赛队伍选边
type MatchTeamSide struct {
	Meta
	Side     string `json:"side"`
	TeamName string `json:"team_name"`
}

// MatchPauseEnable 比赛发起暂停
type MatchPauseEnable struct {
	Meta
	Way string `json:"way"`
}

// MatchPauseDisable 比赛解除暂停
type MatchPauseDisable struct {
	Meta
	Way string `json:"way"`
}

// PlayerConnected 选手连接入游戏
type PlayerConnected struct {
	Meta
	Player  Player `json:"player"`
	Address string `json:"address"`
}

// PlayerDisconnected 选手重连游戏(断开的原因)
type PlayerDisconnected struct {
	Meta
	Player Player `json:"player"`
	Reason string `json:"reason"`
}

// PlayerEntered 选手加入游戏
type PlayerEntered struct {
	Meta
	Player Player `json:"player"`
}

// 选手被服务器禁止
type PlayerBanned struct {
	Meta
	Player   Player `json:"player"`
	Duration string `json:"duration"`
	By       string `json:"by"`
}

// PlayerSwitched 选手换边
type PlayerSwitched struct {
	Meta
	Player Player `json:"player"`
	From   string `json:"from"`
	To     string `json:"to"`
}

// PlayerSay 选手聊天信息
type PlayerSay struct {
	Meta
	Player Player `json:"player"`
	Text   string `json:"text"`
	Team   bool   `json:"team"`
}

// PlayerPurchase 选手购买物品
type PlayerPurchase struct {
	Meta
	Player Player `json:"player"`
	Item   string `json:"item"`
}

// PlayerKill 选手击杀他人
type PlayerKill struct {
	Meta
	Attacker         Player   `json:"attacker"`
	AttackerPosition Position `json:"attacker_pos"`
	Victim           Player   `json:"victim"`
	VictimPosition   Position `json:"victim_pos"`
	Weapon           string   `json:"weapon"`
	Headshot         bool     `json:"headshot"`
	Penetrated       bool     `json:"penetrated"`
	ThroughSmoke     bool     `json:"through_smoke"`
}

// PlayerKillOther 选手击杀地图资源物
type PlayerKillOther struct {
	Meta
	Attacker         Player   `json:"attacker"`
	AttackerPosition Position `json:"attacker_pos"`
	Victim           string   `json:"victim"`
	VictimID         string   `json:"victim_id"`
	VictimPosition   Position `json:"victim_pos"`
	Weapon           string   `json:"weapon"`
}

// PlayerKillAssist 选手击杀助攻(协助队友击杀他人)
type PlayerKillAssist struct {
	Meta
	Attacker Player `json:"attacker"`
	Victim   Player `json:"victim"`
	Flash    bool   `json:"flash"`
}

// PlayerAttack 选手攻击他人
type PlayerAttack struct {
	Meta
	Attacker         Player   `json:"attacker"`
	AttackerPosition Position `json:"attacker_pos"`
	Victim           Player   `json:"victim"`
	VictimPosition   Position `json:"victim_pos"`
	Weapon           string   `json:"weapon"`       // 武器
	Damage           int      `json:"damage"`       // 伤害
	DamageArmor      int      `json:"damage_armor"` // 伤害盔甲
	Health           int      `json:"health"`       // 血量
	Armor            int      `json:"armor"`        // 盔甲
	HitGroup         string   `json:"hitgroup"`     // 身体部位
}

// PlayerKilledBomb 选手被炸弹炸死
type PlayerKilledBomb struct {
	Meta
	Player   Player   `json:"player"`
	Position Position `json:"pos"`
}

// PlayerKilledSuicide 选手自杀
type PlayerKilledSuicide struct {
	Meta
	Player   Player   `json:"player"`
	Position Position `json:"pos"`
	With     string   `json:"with"`
}

// PlayerPickedUp 选手捡起物品
type PlayerPickedUp struct {
	Meta
	Player Player `json:"player"`
	Item   string `json:"item"`
}

// PlayerDropped 选手丢掉物品
type PlayerDropped struct {
	Meta
	Player Player `json:"player"`
	Item   string `json:"item"`
}

// PlayerMoneyChange 选手经济变化
// TODO: add before +-money
type PlayerMoneyChange struct {
	Meta
	Player   Player   `json:"player"`
	Equation Equation `json:"equation"`
	Purchase string   `json:"purchase"`
}

// PlayerBombGot 选手捡起炸弹
type PlayerBombGot struct {
	Meta
	Player Player `json:"player"`
}

// PlayerBombPlanted 选手安放炸弹
type PlayerBombPlanted struct {
	Meta
	Player Player `json:"player"`
	Site   string `json:"site"`
}

// PlayerBombDropped 选手丢掉炸弹
type PlayerBombDropped struct {
	Meta
	Player Player `json:"player"`
}

// PlayerBombBeginDefuse 选手开始拆除炸弹
type PlayerBombBeginDefuse struct {
	Meta
	Player Player `json:"player"`
	Kit    bool   `json:"kit"`
}

// PlayerBombDefused 选手成功拆除炸弹
type PlayerBombDefused struct {
	Meta
	Player Player `json:"player"`
}

// PlayerThrew 选手扔投掷物
type PlayerThrew struct {
	Meta
	Player   Player   `json:"player"`
	Position Position `json:"pos"`
	Entindex int      `json:"entindex"`
	Grenade  string   `json:"grenade"`
}

// PlayerBlinded 选手被闪光
type PlayerBlinded struct {
	Meta
	Attacker Player  `json:"attacker"`
	Victim   Player  `json:"victim"`
	For      float32 `json:"for"`
	Entindex int     `json:"entindex"`
}

// PlayerLeftBuyZone 选手离开购买区域
type PlayerLeftBuyZone struct {
	Meta
	Player Player   `json:"player"`
	Items  []string `json:"items"`
}

// ProjectileSpawned 燃烧弹扩散
type ProjectileSpawned struct {
	Meta
	Position PositionFloat `json:"pos"`
	Velocity Velocity      `json:"velocity"`
}

// ServerCvar 服务器插件
type ServerCvar struct {
	Meta
	Key   string `json:"mode"`
	Value string `json:"value"`
}
