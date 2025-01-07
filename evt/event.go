package evt

import "time"

// 事件数据类型
const (
	UnknownType               = `unknown`
	LogFileStartedType        = `logFileStarted`
	LogFileClosedType         = `logFileClosed`
	LoadingMapType            = `loadingMap`
	ServerMessageType         = `serverMessage`
	ServerCvarType            = `serverCvar`
	RconEventType             = `RconEvent`
	FreezeTimeStartType       = `freezeTimeStart`
	WorldMatchStartType       = `worldMatchStart`
	WorldMatchReloadedType    = `worldReloadedStart`
	WorldRoundStartType       = `worldRoundStart`
	WorldWarmupStartType      = `worldWarmupStart`
	WorldWarmupEndType        = `worldWarmupEnd`
	WorldRoundRestartType     = `worldRoundRestart`
	WorldRoundEndType         = `worldRoundEnd`
	WorldGameCommencingType   = `worldGameCommencing`
	WorldGameOverType         = `worldGameOver`
	WorldKnifeStartType       = `worldGameKnifeStart`
	WorldKnifeWinnerType      = `worldGameKnifeWinner`
	WorldKnifeChoiceType      = `worldGameKnifeChoice`
	MatchTeamScoreType        = `matchTeamScore`
	MatchTeamSideType         = `matchTeamSide`
	MatchPauseEnableType      = `matchPauseEnable`
	MatchPauseDisableType     = `matchPauseDisable`
	TeamScoredType            = `teamScored`
	TeamNoticeType            = `teamNotice`
	PlayerConnectedType       = `playerConnected`
	PlayerDisconnectedType    = `playerDisconnected`
	PlayerEnteredType         = `playerEntered`
	PlayerBannedType          = `playerBanned`
	PlayerSwitchedType        = `playerSwitched`
	PlayerSayType             = `playerSay`
	PlayerPurchaseType        = `playerPurchase`
	PlayerKillType            = `playerKill`
	PlayerKillAssistType      = `playerKillAssist`
	PlayerAttackType          = `playerAttack`
	PlayerKilledBombType      = `playerKilledBomb`
	PlayerKilledSuicideType   = `playerKilledSuicide`
	PlayerKillOtherType       = `playerKillOther`
	PlayerPickedUpType        = `playerPickedUp`
	PlayerDroppedType         = `playerDropped`
	PlayerMoneyChangeType     = `playerMoneyChange`
	PlayerBombGotType         = `playerBombGot`
	PlayerBombPlantedType     = `playerBombPlanted`
	PlayerBombDroppedType     = `playerBombDropped`
	PlayerBombBeginDefuseType = `playerBombBeginDefuse`
	PlayerBombDefusedType     = `playerBombDefused`
	PlayerThrewType           = `playerThrewType`
	PlayerBlindedType         = `playerBlinded`
	ProjectileSpawnedType     = `projectileSpawned`
	PlayerLeftBuyZoneType     = `PlayerLeftBuyZone`
)

// Event 事件类型结构体
type Event struct {
	TypeName string
	Pattern  string
	Func     func(tp string, ti time.Time, r []string) LogMessage
}

type EventList []Event

// GetEventRegister 获取所有支持的事件列表
func GetEventRegister() EventList {
	return registerEventList
}

// 所有注册的日志类型
var registerEventList = []Event{
	{
		TypeName: LogFileStartedType,
		Pattern:  LogFileStartedPattern,
		Func:     NewLogFileStarted,
	},
	{
		TypeName: LogFileClosedType,
		Pattern:  LogFileClosedPattern,
		Func:     NewLogFileClosed,
	},
	{
		TypeName: LoadingMapType,
		Pattern:  LoadingMapPattern,
		Func:     NewLoadingMap,
	},
	{
		TypeName: ServerMessageType,
		Pattern:  ServerMessagePattern,
		Func:     NewServerMessage,
	},
	{
		TypeName: ServerCvarType,
		Pattern:  ServerCvarPattern,
		Func:     NewServerCvar,
	},

	{
		TypeName: RconEventType,
		Pattern:  RconEventPattern,
		Func:     NewRconEvent,
	},
	{
		TypeName: FreezeTimeStartType,
		Pattern:  FreezeTimeStartPattern,
		Func:     NewFreezeTimeStart,
	},
	{
		TypeName: WorldMatchStartType,
		Pattern:  WorldMatchStartPattern,
		Func:     NewWorldMatchStart,
	},
	{
		TypeName: WorldMatchReloadedType,
		Pattern:  WorldMatchReloadedPattern,
		Func:     NewWorldReloadedStart,
	},
	{
		TypeName: WorldWarmupStartType,
		Pattern:  WorldWarmupStartPattern,
		Func:     NewWorldWarmupStart,
	},
	{
		TypeName: WorldWarmupEndType,
		Pattern:  WorldWarmupEndPattern,
		Func:     NewWorldWarmupEnd,
	},
	{
		TypeName: WorldRoundStartType,
		Pattern:  WorldRoundStartPattern,
		Func:     NewWorldRoundStart,
	},
	{
		TypeName: WorldRoundRestartType,
		Pattern:  WorldRoundRestartPattern,
		Func:     NewWorldRoundRestart,
	},
	{
		TypeName: WorldRoundEndType,
		Pattern:  WorldRoundEndPattern,
		Func:     NewWorldRoundEnd,
	},
	{
		TypeName: WorldGameCommencingType,
		Pattern:  WorldGameCommencingPattern,
		Func:     NewWorldGameCommencing,
	},
	{
		TypeName: WorldKnifeStartType,
		Pattern:  WorldKnifeStartPattern,
		Func:     NewWorldKnifeStart,
	},
	{
		TypeName: WorldKnifeWinnerType,
		Pattern:  WorldKnifeWinnerPattern,
		Func:     NewWorldKnifeWinner,
	},
	{
		TypeName: WorldKnifeChoiceType,
		Pattern:  WorldKnifeChoicePattern,
		Func:     NewWorldKnifeChoice,
	},
	{
		TypeName: MatchTeamScoreType,
		Pattern:  MatchTeamScorePattern,
		Func:     NewMatchTeamScore,
	},
	{
		TypeName: MatchTeamSideType,
		Pattern:  MatchTeamSidePattern,
		Func:     NewMatchTeamSide,
	},
	{
		TypeName: MatchPauseEnableType,
		Pattern:  MatchPauseEnablePattern,
		Func:     NewMatchPauseEnable,
	},
	{
		TypeName: MatchPauseDisableType,
		Pattern:  MatchPauseDisablePattern,
		Func:     NewMatchPauseDisable,
	},
	{
		TypeName: MatchTeamSideType,
		Pattern:  MatchTeamSidePattern,
		Func:     NewMatchTeamSide,
	},
	{
		TypeName: TeamScoredType,
		Pattern:  TeamScoredPattern,
		Func:     NewTeamScored,
	},
	{
		TypeName: TeamNoticeType,
		Pattern:  TeamNoticePattern,
		Func:     NewTeamNotice,
	},
	{
		TypeName: PlayerConnectedType,
		Pattern:  PlayerConnectedPattern,
		Func:     NewPlayerConnected,
	},
	{
		TypeName: PlayerDisconnectedType,
		Pattern:  PlayerDisconnectedPattern,
		Func:     NewPlayerDisconnected,
	},
	{
		TypeName: PlayerEnteredType,
		Pattern:  PlayerEnteredPattern,
		Func:     NewPlayerEntered,
	},
	{
		TypeName: PlayerBannedType,
		Pattern:  PlayerBannedPattern,
		Func:     NewPlayerBanned,
	},
	{
		TypeName: PlayerSwitchedType,
		Pattern:  PlayerSwitchedPattern,
		Func:     NewPlayerSwitched,
	},
	{
		TypeName: PlayerSayType,
		Pattern:  PlayerSayPattern,
		Func:     NewPlayerSay,
	},
	{
		TypeName: PlayerPurchaseType,
		Pattern:  PlayerPurchasePattern,
		Func:     NewPlayerPurchase,
	},
	{
		TypeName: PlayerKillType,
		Pattern:  PlayerKillPattern,
		Func:     NewPlayerKill,
	},
	{
		TypeName: PlayerKillOtherType,
		Pattern:  PlayerKillOtherPattern,
		Func:     NewPlayerKillOther,
	},
	{
		TypeName: PlayerKillAssistType,
		Pattern:  PlayerKillAssistPattern,
		Func:     NewPlayerKillAssist,
	},
	{
		TypeName: PlayerAttackType,
		Pattern:  PlayerAttackPattern,
		Func:     NewPlayerAttack,
	},
	{
		TypeName: PlayerKilledBombType,
		Pattern:  PlayerKilledBombPattern,
		Func:     NewPlayerKilledBomb,
	},
	{
		TypeName: PlayerKilledSuicideType,
		Pattern:  PlayerKilledSuicidePattern,
		Func:     NewPlayerKilledSuicide,
	},
	{
		TypeName: PlayerPickedUpType,
		Pattern:  PlayerPickedUpPattern,
		Func:     NewPlayerPickedUp,
	},
	{
		TypeName: PlayerDroppedType,
		Pattern:  PlayerDroppedPattern,
		Func:     NewPlayerDropped,
	},
	{
		TypeName: PlayerMoneyChangeType,
		Pattern:  PlayerMoneyChangePattern,
		Func:     NewPlayerMoneyChange,
	},
	{
		TypeName: PlayerBombGotType,
		Pattern:  PlayerBombGotPattern,
		Func:     NewPlayerBombGot,
	},
	{
		TypeName: PlayerBombPlantedType,
		Pattern:  PlayerBombPlantedPattern,
		Func:     NewPlayerBombPlanted,
	},
	{
		TypeName: PlayerBombDroppedType,
		Pattern:  PlayerBombDroppedPattern,
		Func:     NewPlayerBombDropped,
	},
	{
		TypeName: PlayerBombBeginDefuseType,
		Pattern:  PlayerBombBeginDefusePattern,
		Func:     NewPlayerBombBeginDefuse,
	},
	{
		TypeName: PlayerBombDefusedType,
		Pattern:  PlayerBombDefusedPattern,
		Func:     NewPlayerBombDefused,
	},
	{
		TypeName: PlayerThrewType,
		Pattern:  PlayerThrewPattern,
		Func:     NewPlayerThrew,
	},
	{
		TypeName: PlayerBlindedType,
		Pattern:  PlayerBlindedPattern,
		Func:     NewPlayerBlinded,
	}, {
		TypeName: PlayerLeftBuyZoneType,
		Pattern:  PlayerLeftBuyZonePattern,
		Func:     NewPlayerLeftBuyZone,
	},
	{
		TypeName: ProjectileSpawnedType,
		Pattern:  ProjectileSpawnedPattern,
		Func:     NewProjectileSpawned,
	},
	{
		TypeName: WorldGameOverType,
		Pattern:  WorldGameOverPattern,
		Func:     NewWorldGameOver,
	},
}

// AppendRegisterEvent 自定义附加注册事件
func AppendRegisterEvent(e Event) {
	registerEventList = append(registerEventList, e)
}
