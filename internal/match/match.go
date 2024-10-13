package match

import "github.com/kikudesuyo/gopher-eleven/internal/team"

type Turn struct {
	count       int
	offenceTeam *team.Team
}

type Period string

const (
	firstHalf  Period = "FIRST_HALF"
	secondHalf Period = "SECOND_HALF"
	overTime   Period = "OVER_TIME"
)

type Match struct {
	turn         Turn
	period       Period
	playerTeam   team.Team
	opponentTeam team.Team
}

func InitMatch() Match {
	playerTeam := team.GetPlayerTeam()
	opponentTeam := team.GetOpponentTeam()
	match := Match{
		turn:         Turn{0, &playerTeam},
		period:       firstHalf,
		playerTeam:   playerTeam,
		opponentTeam: opponentTeam,
	}
	return match
}
