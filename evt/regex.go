package evt

// 日志类型正则表达式
const (
	// LogLineValidPattern 行日志的数据头标识
	LogLineValidPattern = `L (\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}): (.*)`
	// HttpLineValidPattern HTTP行日志的数据头标识
	HttpLineValidPattern = `(\d{2}\/\d{2}\/\d{4} - \d{2}:\d{2}:\d{2}.\d{3}) - (.*)`
)

// 事件数据正则表达式
const (
	//LogFileStartedPattern Log file started log文件开始
	LogFileStartedPattern = `Log file started (file "(.+)") (game "(\w+)") (version "(\w+)")`
	LogFileClosedPattern  = `Log file closed`
	//LoadingMapPattern 加载地图
	LoadingMapPattern = `Loading map "(\w+)"`
	// ServerMessagePattern 服务器信息
	ServerMessagePattern = `server_message: "(\w+)"`
	// FreezeTimeStartPattern 冻结时间开始
	FreezeTimeStartPattern = `Starting Freeze period`
	// WorldMatchStartPattern 比赛开始
	WorldMatchStartPattern = `World triggered "Match_Start" on "(\w+)"`
	// WorldMatchReloadedPattern 比赛回档
	WorldMatchReloadedPattern = `World triggered "Match_Reloaded" on "(\w+)"`
	//WorldWarmupStartPattern 热身开始
	WorldWarmupStartPattern = `World triggered "Warmup_Start"`
	//WorldWarmupEndPattern 热身结束
	WorldWarmupEndPattern = `World triggered "Warmup_End"`
	// WorldRoundStartPattern 回合开始
	WorldRoundStartPattern = `World triggered "Round_Start"`
	// WorldRoundRestartPattern 回合重启
	WorldRoundRestartPattern = `World triggered "Restart_Round_\((\d+)_second\)`
	// WorldRoundEndPattern 回合结束
	WorldRoundEndPattern = `World triggered "Round_End"`
	// WorldGameCommencingPattern 游戏开始
	WorldGameCommencingPattern = `World triggered "Game_Commencing"`
	// WorldKnifeStartPattern 刀局开始
	WorldKnifeStartPattern = `World triggered "Knife_Round_Start"`
	// WorldKnifeWinnerPattern 刀局胜利
	WorldKnifeWinnerPattern = `World triggered "Knife_Round_Winner" (\w+)`
	// WorldKnifeChoicePattern 刀局胜方选择结果
	WorldKnifeChoicePattern = `World triggered "Knife_Round_Choice" (\w+)`
	// WorldGameOverPattern 游戏结束
	WorldGameOverPattern = `Game Over: (\w+) (\w+) (\w+) score (\d+):(\d+) after (\d+) min`
	// MatchTeamScorePattern 比赛队伍比分
	MatchTeamScorePattern = `MatchStatus: Score: (\d+):(\d+) on map \"(.+)\" RoundsPlayed: (-?\d+)`
	// MatchTeamSidePattern 比赛队伍选边
	MatchTeamSidePattern = `(MatchStatus: )?Team playing "(CT|TERRORIST)": (.*)`
	// MatchPauseEnablePattern 比赛发起暂停
	MatchPauseEnablePattern = `Match pause is enabled - (\w+)`
	// MatchPauseDisablePattern 比赛暂停解除
	MatchPauseDisablePattern = `Match pause is disabled - (\w+)`
	// TeamScoredPattern 队伍得分
	TeamScoredPattern = `Team "(CT|TERRORIST)" scored "(\d+)" with "(\d+)" players`
	// TeamNoticePattern 队伍比分结果
	TeamNoticePattern = `Team "(CT|TERRORIST)" triggered "(\w+)" \(CT "(\d+)"\) \(T "(\d+)"\)`
	// PlayerConnectedPattern 选手链接
	PlayerConnectedPattern = `"(.+)<(\d+)><(.+)><>" connected, address "(.*)"`
	// PlayerDisconnectedPattern 选手断开链接
	PlayerDisconnectedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT|Unassigned|Spectator)>" disconnected \(reason "(.+)"\)`
	// PlayerEnteredPattern 选手加入游戏
	PlayerEnteredPattern = `"(.+)<(\d+)><(.+)><>" entered the game`
	// PlayerBannedPattern 选手被banned
	PlayerBannedPattern = `Banid: "(.+)<(\d+)><(.+)><\w*>" was banned "([\w. ]+)" by "(\w+)"`
	// PlayerSwitchedPattern 选手换边
	PlayerSwitchedPattern = `"(.+)<(\d+)><(.+)>" switched from team <(Unassigned|Spectator|TERRORIST|CT)> to <(Unassigned|Spectator|TERRORIST|CT)>`
	// PlayerSayPattern 选手交流
	PlayerSayPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" say(_team)? "(.*)"`
	// PlayerPurchasePattern 选手购买物品
	PlayerPurchasePattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" purchased "(\w+)"`
	// PlayerKillPattern 选手击杀
	PlayerKillPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] killed "(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] with "(\w+)" ?(\(?(headshot|penetrated|throughsmoke|headshot penetrated|headshot throughsmoke)?\))?`
	// PlayerKillAssistPattern 选手助攻击杀（assisted/flash-assisted）
	PlayerKillAssistPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" (.+) killing "(.+)<(\d+)><(.+)><(TERRORIST|CT)>"`
	// PlayerAttackPattern 选手攻击
	PlayerAttackPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] attacked "(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] with "(\w+)" \(damage "(\d+)"\) \(damage_armor "(\d+)"\) \(health "(\d+)"\) \(armor "(\d+)"\) \(hitgroup "([\w ]+)"\)`
	// PlayerKilledBombPattern 选手被炸弹击杀
	PlayerKilledBombPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] was killed by the bomb\.`
	// PlayerKilledSuicidePattern 选手自杀
	PlayerKilledSuicidePattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] committed suicide with "(.*)"`
	// PlayerKillOtherPattern 选手击杀地图资源物
	PlayerKillOtherPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" \[(-?\d+) (-?\d+) (-?\d+)\] killed other "(.+)<(\d+)>" \[(-?\d+) (-?\d+) (-?\d+)\] with "(\w+)"`
	// PlayerPickedUpPattern 选手捡起物品
	PlayerPickedUpPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" picked up "(\w+)"`
	// PlayerDroppedPattern 选手丢掉物品
	PlayerDroppedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT|Unassigned)>" dropped "(\w+)"`
	// PlayerMoneyChangePattern 选手经济变化
	PlayerMoneyChangePattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" money change (\d+)\+?(-?\d+) = \$(\d+)( \(tracked\))?( \(purchase: (\w+)\))?`
	// PlayerBombGotPattern 选手获得炸弹
	PlayerBombGotPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" triggered "Got_The_Bomb"`
	// PlayerBombPlantedPattern 选手安放炸弹
	PlayerBombPlantedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" triggered "Planted_The_Bomb" at bombsite (\w+)`
	// PlayerBombDroppedPattern 选手丢掉炸弹
	PlayerBombDroppedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" triggered "Dropped_The_Bomb"`
	// PlayerBombBeginDefusePattern 选手开始解除炸弹
	PlayerBombBeginDefusePattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" triggered "Begin_Bomb_Defuse_With(out)?_Kit"`
	// PlayerBombDefusedPattern 选手成功解除炸弹
	PlayerBombDefusedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" triggered "Defused_The_Bomb"`
	// PlayerThrewPattern 选手扔投掷物
	PlayerThrewPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" threw (\w+) \[(-?\d+) (-?\d+) (-?\d+)\]( flashbang entindex (\d+))?\)?`
	// PlayerBlindedPattern 选手被闪光
	PlayerBlindedPattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" blinded for ([\d.]+) by "(.+)<(\d+)><(.+)><(TERRORIST|CT)>" from flashbang entindex (\d+)`
	// PlayerLeftBuyZonePattern 选手离开购买区域
	PlayerLeftBuyZonePattern = `"(.+)<(\d+)><(.+)><(TERRORIST|CT)>" left buyzone with \[(.*)\]`
	// ProjectileSpawnedPattern regular expression
	ProjectileSpawnedPattern = `Molotov projectile spawned at (-?\d+\.\d+) (-?\d+\.\d+) (-?\d+\.\d+), velocity (-?\d+\.\d+) (-?\d+\.\d+) (-?\d+\.\d+)`
	// ServerCvarPattern regular expression
	ServerCvarPattern = `server_cvar: "(\w+)" "(.*)"`
	// RconEventPattern regular expression
	RconEventPattern = `rcon from "(.*):(\d+)": command "(.*)"`

	//TODO:  World triggered "SFUI_Notice_Round_Draw" (CT "0") (T "1")
	// TODO // VoteStartPattern = `Vote started "StartTimeOut " from #2 "416<16><STEAM_1:1:55894410><TERRORIST><Area 4>"`
	// TODO // VoteCastPattern = `Vote cast "StartTimeOut " from #2 "416<16><STEAM_1:1:55894410><TERRORIST><Area 4>" option0`
	// TODO // VoteSuccessPattern = `Vote cast "StartTimeOut " from #2 "416<16><STEAM_1:1:55894410><TERRORIST><Area 4>`
)
