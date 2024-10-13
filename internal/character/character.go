package character

import "github.com/kikudesuyo/gopher-eleven/internal"

type Technique struct {
	attr  string
	cost  int
	name  string
	power int
}

type Character struct {
	Name       string
	Techniques []Technique
	Tp         int
}

func GetPlayerTeamCharacters() []Character {
	ptids := internal.GetPlayerTeamCharacterIds()
	playerTeamCharacters := make([]Character, 2)

	for idx, id := range ptids {
		character := internal.GetCharacter(id)
		ids := character["techniques"].([]string)
		techniques := make([]Technique, len(ids))

		for idx, id := range ids {
			t := internal.GetTechniques(id)
			techniques[idx] = Technique{
				attr:  t["attr"].(string),
				cost:  t["cost"].(int),
				name:  t["name"].(string),
				power: t["power"].(int),
			}
		}
		playerTeamCharacters[idx] = Character{
			Name:       character["name"].(string),
			Techniques: techniques,
			Tp:         character["tp"].(int),
		}
	}
	return playerTeamCharacters
}

func GetOpponentTeamCharacters() []Character {
	opids := internal.GetOpponentTeamCharacterIds()
	opponentTeamCharacters := make([]Character, 2)
	for idx, id := range opids {
		character := internal.GetCharacter(id)
		ids := character["techniques"].([]string)
		techniques := make([]Technique, len(ids))
		for idx, id := range ids {
			t := internal.GetTechniques(id)
			techniques[idx] = Technique{
				attr:  t["attr"].(string),
				cost:  t["cost"].(int),
				name:  t["name"].(string),
				power: t["power"].(int),
			}
		}
		opponentTeamCharacters[idx] = Character{
			Name:       character["name"].(string),
			Techniques: techniques,
			Tp:         character["tp"].(int),
		}
	}
	return opponentTeamCharacters
}