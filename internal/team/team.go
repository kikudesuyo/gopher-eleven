package team

import (
	"github.com/kikudesuyo/gopher-eleven/internal/character"
	"github.com/kikudesuyo/gopher-eleven/internal/db"
)

type Team struct {
	Characters []character.Character
	Name       string
	point      int
}

func GetPlayerTeam() Team {
	return Team{
		Characters: character.GetPlayerTeamCharacters(),
		Name:       db.GetPlayerTeamName(),
		point:      0,
	}
}

func GetOpponentTeam() Team {
	return Team{
		Characters: character.GetOpponentTeamCharacters(),
		Name:       db.GetOpponentTeamName(),
		point:      0,
	}
}
