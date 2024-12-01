package evt

import (
	"strconv"
	"strings"
	"time"
)

type LogMessage interface {
	GetType() string
	GetTime() time.Time
}

// ------------------------------------------------------------------------------------------------------------------------

func NewUnknown(ti time.Time, r []string) LogMessage {
	return Unknown{
		Meta: NewMeta(ti, UnknownType),
		Raw:  r[1],
	}
}
func NewServerMessage(ty string, ti time.Time, r []string) LogMessage {
	return ServerMessage{
		Meta: NewMeta(ti, ty),
		Text: r[1],
	}
}

func NewFreezeTimeStart(ty string, ti time.Time, r []string) LogMessage {
	return FreezeTimeStart{NewMeta(ti, ty)}
}

func NewWorldMatchStart(ty string, ti time.Time, r []string) LogMessage {
	return WorldMatchStart{
		Meta: NewMeta(ti, ty),
		Map:  r[1],
	}
}

func NewWorldReloadedStart(ty string, ti time.Time, r []string) LogMessage {
	return WorldReloadedStart{
		Meta: NewMeta(ti, ty),
		Map:  r[1],
	}
}

func NewWorldRoundStart(ty string, ti time.Time, r []string) LogMessage {
	return WorldRoundStart{NewMeta(ti, ty)}
}

func NewWorldRoundRestart(ty string, ti time.Time, r []string) LogMessage {
	return WorldRoundRestart{
		Meta:     NewMeta(ti, ty),
		TimeLeft: toInt(r[1]),
	}
}

func NewWorldRoundEnd(ty string, ti time.Time, r []string) LogMessage {
	return WorldRoundEnd{NewMeta(ti, ty)}
}

func NewWorldGameCommencing(ty string, ti time.Time, r []string) LogMessage {
	return WorldGameCommencing{NewMeta(ti, ty)}
}

func NewWorldKnifeStart(ty string, ti time.Time, r []string) LogMessage {
	return WorldKnifeStart{NewMeta(ti, ty)}
}

func NewWorldKnifeWinner(ty string, ti time.Time, r []string) LogMessage {
	return WorldKnifeWinner{Meta: NewMeta(ti, ty), Side: r[1]}
}

func NewWorldKnifeChoice(ty string, ti time.Time, r []string) LogMessage {
	return WorldKnifeChoice{Meta: NewMeta(ti, ty), Side: r[1]}
}

func NewMatchTeamScore(ty string, ti time.Time, r []string) LogMessage {
	return MatchTeamScore{
		Meta:         NewMeta(ti, ty),
		ScoreCT:      toInt(r[1]),
		ScoreT:       toInt(r[2]),
		Map:          r[3],
		RoundsPlayed: toInt(r[4]),
	}
}

func NewMatchTeamSide(ty string, ti time.Time, r []string) LogMessage {
	return MatchTeamSide{
		Meta:     NewMeta(ti, ty),
		Side:     r[2],
		TeamName: r[3],
	}
}

func NewMatchPauseEnable(ty string, ti time.Time, r []string) LogMessage {
	return MatchPauseEnable{
		Meta: NewMeta(ti, ty),
		Way:  r[1],
	}
}

func NewMatchPauseDisable(ty string, ti time.Time, r []string) LogMessage {
	return MatchPauseDisable{
		Meta: NewMeta(ti, ty),
		Way:  r[1],
	}
}

func NewTeamScored(ty string, ti time.Time, r []string) LogMessage {
	return TeamScored{
		Meta:       NewMeta(ti, ty),
		Side:       r[1],
		Score:      toInt(r[2]),
		NumPlayers: toInt(r[3]),
	}
}

func NewTeamNotice(ty string, ti time.Time, r []string) LogMessage {
	return TeamNotice{
		Meta:    NewMeta(ti, ty),
		Side:    r[1],
		Notice:  r[2],
		ScoreCT: toInt(r[3]),
		ScoreT:  toInt(r[4]),
	}
}

func NewPlayerConnected(ty string, ti time.Time, r []string) LogMessage {
	return PlayerConnected{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    "",
		},
		Address: r[4],
	}
}

func NewPlayerDisconnected(ty string, ti time.Time, r []string) LogMessage {
	return PlayerDisconnected{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Reason: r[5],
	}
}

func NewPlayerEntered(ty string, ti time.Time, r []string) LogMessage {
	return PlayerEntered{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    "",
		},
	}
}

func NewPlayerBanned(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBanned{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    "",
		},
		Duration: r[4],
		By:       r[5],
	}
}

func NewPlayerSwitched(ty string, ti time.Time, r []string) LogMessage {
	return PlayerSwitched{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    "",
		},
		From: r[4],
		To:   r[5],
	}
}

func NewPlayerSay(ty string, ti time.Time, r []string) LogMessage {
	return PlayerSay{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Team: r[5] == "_team",
		Text: r[6],
	}
}

func NewPlayerPurchase(ty string, ti time.Time, r []string) LogMessage {
	return PlayerPurchase{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Item: r[5],
	}
}

func NewPlayerKill(ty string, ti time.Time, r []string) LogMessage {
	return PlayerKill{
		Meta: NewMeta(ti, ty),
		Attacker: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		AttackerPosition: Position{
			X: toInt(r[5]),
			Y: toInt(r[6]),
			Z: toInt(r[7]),
		},
		Victim: Player{
			Name:    r[8],
			ID:      toInt(r[9]),
			SteamID: r[10],
			Side:    r[11],
		},
		VictimPosition: Position{
			X: toInt(r[12]),
			Y: toInt(r[13]),
			Z: toInt(r[14]),
		},
		Weapon:       r[15],
		Headshot:     strings.Contains(r[17], "headshot"),
		Penetrated:   strings.Contains(r[17], "penetrated"),
		ThroughSmoke: strings.Contains(r[17], "throughsmoke"),
	}
}

func NewPlayerKillAssist(ty string, ti time.Time, r []string) LogMessage {
	return PlayerKillAssist{
		Meta: NewMeta(ti, ty),
		Attacker: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Victim: Player{
			Name:    r[6],
			ID:      toInt(r[7]),
			SteamID: r[8],
			Side:    r[9],
		},
		Flash: strings.Contains(r[5], "flash"),
	}
}

func NewPlayerAttack(ty string, ti time.Time, r []string) LogMessage {
	return PlayerAttack{
		Meta: NewMeta(ti, ty),
		Attacker: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		AttackerPosition: Position{
			X: toInt(r[5]),
			Y: toInt(r[6]),
			Z: toInt(r[7]),
		},
		Victim: Player{
			Name:    r[8],
			ID:      toInt(r[9]),
			SteamID: r[10],
			Side:    r[11],
		},
		VictimPosition: Position{
			X: toInt(r[12]),
			Y: toInt(r[13]),
			Z: toInt(r[14]),
		},
		Weapon:      r[15],
		Damage:      toInt(r[16]),
		DamageArmor: toInt(r[17]),
		Health:      toInt(r[18]),
		Armor:       toInt(r[19]),
		HitGroup:    r[20],
	}
}

func NewPlayerKilledBomb(ty string, ti time.Time, r []string) LogMessage {
	return PlayerKilledBomb{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Position: Position{
			X: toInt(r[5]),
			Y: toInt(r[6]),
			Z: toInt(r[7]),
		},
	}
}

func NewPlayerKilledSuicide(ty string, ti time.Time, r []string) LogMessage {
	return PlayerKilledSuicide{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Position: Position{
			X: toInt(r[5]),
			Y: toInt(r[6]),
			Z: toInt(r[7]),
		},
		With: r[8],
	}
}

func NewPlayerPickedUp(ty string, ti time.Time, r []string) LogMessage {
	return PlayerPickedUp{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Item: r[5],
	}
}

func NewPlayerDropped(ty string, ti time.Time, r []string) LogMessage {
	return PlayerDropped{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Item: r[5],
	}
}

func NewPlayerMoneyChange(ty string, ti time.Time, r []string) LogMessage {
	return PlayerMoneyChange{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Equation: Equation{
			A:      toInt(r[5]),
			B:      toInt(r[6]),
			Result: toInt(r[7]),
		},
		Purchase: r[10],
	}
}

func NewPlayerBombGot(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBombGot{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
	}
}

func NewPlayerBombPlanted(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBombPlanted{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Site: r[5],
	}
}

func NewPlayerBombDropped(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBombDropped{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
	}
}

func NewPlayerBombBeginDefuse(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBombBeginDefuse{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Kit: !(r[5] == "out"),
	}
}

func NewPlayerBombDefused(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBombDefused{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
	}
}

func NewPlayerThrew(ty string, ti time.Time, r []string) LogMessage {
	return PlayerThrew{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Grenade: r[5],
		Position: Position{
			X: toInt(r[6]),
			Y: toInt(r[7]),
			Z: toInt(r[8]),
		},
		Entindex: toInt(r[10]),
	}
}

func NewPlayerBlinded(ty string, ti time.Time, r []string) LogMessage {
	return PlayerBlinded{
		Meta: NewMeta(ti, ty),
		Victim: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		For: toFloat32(r[5]),
		Attacker: Player{
			Name:    r[6],
			ID:      toInt(r[7]),
			SteamID: r[8],
			Side:    r[9],
		},
		Entindex: toInt(r[10]),
	}
}

func NewPlayerLeftBuyZone(ty string, ti time.Time, r []string) LogMessage {
	return PlayerLeftBuyZone{
		Meta: NewMeta(ti, ty),
		Player: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		Items: strings.Split(strings.Trim(r[5], " "), " "),
	}
}

func NewProjectileSpawned(ty string, ti time.Time, r []string) LogMessage {
	return ProjectileSpawned{
		Meta: NewMeta(ti, ty),
		Position: PositionFloat{
			X: toFloat32(r[1]),
			Y: toFloat32(r[2]),
			Z: toFloat32(r[3]),
		},
		Velocity: Velocity{
			X: toFloat32(r[4]),
			Y: toFloat32(r[5]),
			Z: toFloat32(r[6]),
		},
	}
}

func NewWorldGameOver(ty string, ti time.Time, r []string) LogMessage {
	return WorldGameOver{
		Meta:     NewMeta(ti, ty),
		Mode:     r[1],
		MapGroup: r[2],
		Map:      r[3],
		ScoreCT:  toInt(r[4]),
		ScoreT:   toInt(r[5]),
		Duration: toInt(r[6]),
	}
}

func NewServerCvar(ty string, ti time.Time, r []string) LogMessage {
	return ServerCvar{
		Meta:  NewMeta(ti, ty),
		Key:   r[1],
		Value: r[2],
	}
}

func NewRconEvent(ty string, ti time.Time, r []string) LogMessage {
	// r[1]=ip r[2]=port r[3]=command
	p, err := strconv.Atoi(r[2])
	if err != nil {
		return NewUnknown(ti, r)
	}
	return Rcon{
		Meta:    NewMeta(ti, ty),
		IP:      r[1],
		Port:    uint(p),
		Command: r[3],
	}
}

func NewPlayerKillOther(ty string, ti time.Time, r []string) LogMessage {
	return PlayerKillOther{
		Meta: NewMeta(ti, ty),
		Attacker: Player{
			Name:    r[1],
			ID:      toInt(r[2]),
			SteamID: r[3],
			Side:    r[4],
		},
		AttackerPosition: Position{
			X: toInt(r[5]),
			Y: toInt(r[6]),
			Z: toInt(r[7]),
		},
		Victim:   r[8],
		VictimID: r[9],
		VictimPosition: Position{
			X: toInt(r[10]),
			Y: toInt(r[11]),
			Z: toInt(r[12]),
		},
		Weapon: r[13],
	}
}

// toInt converts string to int, assigns 0 when not convertable
func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

func toFloat32(v string) float32 {
	i, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return float32(0)
	}
	return float32(i)
}
