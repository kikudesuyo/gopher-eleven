package main

import (
	"github.com/kikudesuyo/gopher-eleven/internal"
)

type Character struct {
	Name              string
	SpecialTechniques []string
	Tp                int
}

func main() {
	ptids := internal.GetPlayerTeamCharacterIds()
	playerTeamCharacters := make([]Character, 2)
	for idx, id := range ptids {
		character := internal.GetCharacter(id)
		playerTeamCharacters[idx] = Character{
			Name:              character["name"].(string),
			SpecialTechniques: character["specialTechniques"].([]string),
			Tp:                character["tp"].(int),
		}
	}
	otids := internal.GetOpponentTeamCharacterIds()
	opponentTeamCharacters := make([]Character, 2)
	for idx, id := range otids {
		character := internal.GetCharacter(id)
		opponentTeamCharacters[idx] = Character{
			Name:              character["name"].(string),
			SpecialTechniques: character["specialTechniques"].([]string),
			Tp:                character["tp"].(int),
		}
	}
}
