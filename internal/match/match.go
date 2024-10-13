package match

import (
	"errors"

	"github.com/kikudesuyo/gopher-eleven/internal/team"
)

type Turn struct {
	count       int
	offenceTeam *team.Team
}

func (t *Turn) Inc() {
	t.count++
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

func (m *Match) Proceed() (string, string, bool, error) {
	offenceTeam := m.turn.offenceTeam
	cn, tn := offenceTeam.Characters[0].Perform()
	m.turn.Inc()

	switch m.turn.offenceTeam.Name {
	case m.playerTeam.Name:
		m.turn.offenceTeam = &m.opponentTeam
	case m.opponentTeam.Name:
		m.turn.offenceTeam = &m.playerTeam
	default:
		return "", "", true, errors.New("invalid team name")
	}

	// switch m.turn.offenceTeam {
	// case &m.playerTeam:
	// 	m.turn.offenceTeam = &m.opponentTeam
	// case &m.opponentTeam:
	// 	m.turn.offenceTeam = &m.playerTeam
	// default:
	// 	return "", "", errors.New("invalid team")
	// }

	isEnd := m.turn.count == 3
	return cn, tn, isEnd, nil
}
