package character

import (
	"math/rand"

	"github.com/kikudesuyo/gopher-eleven/internal/db"
)

type Technique struct {
	Attr  string
	Cost  int
	Name  string
	Power int
}

type Character struct {
	Name       string
	Techniques []Technique
	Tp         int
}

func (c *Character) Perform() Technique {

	idx := rand.Intn(len(c.Techniques))
	c.Tp -= c.Techniques[idx].Cost
	return c.Techniques[idx]
}

func GetPlayerTeamCharacters() []Character {
	ptids := db.GetPlayerTeamCharacterIds()
	playerTeamCharacters := make([]Character, 2)

	for idx, id := range ptids {
		character := db.GetCharacter(id)
		ids := character["techniques"].([]string)
		techniques := make([]Technique, len(ids))

		for idx, id := range ids {
			t := db.GetTechniques(id)
			techniques[idx] = Technique{
				Attr:  t["attr"].(string),
				Cost:  t["cost"].(int),
				Name:  t["name"].(string),
				Power: t["power"].(int),
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
	opids := db.GetOpponentTeamCharacterIds()
	opponentTeamCharacters := make([]Character, 2)
	for idx, id := range opids {
		character := db.GetCharacter(id)
		ids := character["techniques"].([]string)
		techniques := make([]Technique, len(ids))
		for idx, id := range ids {
			t := db.GetTechniques(id)
			techniques[idx] = Technique{
				Attr:  t["attr"].(string),
				Cost:  t["cost"].(int),
				Name:  t["name"].(string),
				Power: t["power"].(int),
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
