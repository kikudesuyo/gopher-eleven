package match

import (
	"github.com/kikudesuyo/gopher-eleven/internal/character"
	"github.com/kikudesuyo/gopher-eleven/internal/team"
)

type Turn struct {
	count       int
	offenceTeam *team.Team
	defenceTeam *team.Team
}

func (t *Turn) incCount() {
	t.count++
}

func (t *Turn) swapOffenceAndDefence() {
	t.offenceTeam, t.defenceTeam = t.defenceTeam, t.offenceTeam
}

func (t *Turn) isEnd() bool {
	return t.count == 4
}

func (t *Turn) performTechnique(texts []string) (character.Character, character.Technique, character.Character, character.Technique, []string) {
	offenceCharacter := t.offenceTeam.Characters[1]
	offenceTechnique, offenceText := offenceCharacter.Perform()
	defenceCharacter := t.defenceTeam.Characters[0]
	defenceTechnique, defenceText := defenceCharacter.Perform()
	texts = append(texts, offenceText, defenceText)
	return offenceCharacter, offenceTechnique, defenceCharacter, defenceTechnique, texts
}

func (t *Turn) appendTurnResultTexts(texts []string, offencePower, defencePower int, offenceCharacterName, defenceCharacterName string) []string {
	if isPowerGreater(offencePower, defencePower) {
		t.offenceTeam.IncScore()
		texts = append(texts, "角間「決まったぁぁーー! "+offenceCharacterName+"のシュートが炸裂!!」")
	} else {
		texts = append(texts, "角間「キーパーの"+defenceCharacterName+"がしっかりキャッチ!」")
	}
	texts = append(texts, "------------------------")
	return texts
}
