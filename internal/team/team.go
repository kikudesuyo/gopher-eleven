package team

import (
	"github.com/kikudesuyo/gopher-eleven/internal/character"
)

type Team struct {
	characters []character.Character
	point      int
}

func GetPlayerTeam() Team {
	return Team{
		characters: character.GetPlayerTeamCharacters(),
		point:      0,
	}
}

func GetOpponentTeam() Team {
	return Team{
		characters: character.GetOpponentTeamCharacters(),
		point:      0,
	}
}
