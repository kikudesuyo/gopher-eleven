package team

import (
	"github.com/kikudesuyo/gopher-eleven/internal/character"
	"github.com/kikudesuyo/gopher-eleven/internal/db"
)

type Team struct {
	Characters []character.Character
	Name       string
	Score      int
}

func GetPlayerTeam() Team {
	return Team{
		Characters: character.GetPlayerTeamCharacters(),
		Name:       db.GetPlayerTeamName(),
		Score:      0,
	}
}

func GetOpponentTeam() Team {
	return Team{
		Characters: character.GetOpponentTeamCharacters(),
		Name:       db.GetOpponentTeamName(),
		Score:      0,
	}
}

func (t *Team) IncScore() int {
	t.Score++
	return t.Score
}
