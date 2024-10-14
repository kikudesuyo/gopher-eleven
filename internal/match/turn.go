package match

import (
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

func (t *Turn) performTechnique(texts []string) []string {
	offenceCharacter := t.offenceTeam.Characters[1]
	offenceTechnique, offenceText := offenceCharacter.Perform()
	defenceCharacter := t.defenceTeam.Characters[0]
	defenceTechnique, defenceText := defenceCharacter.Perform()
	texts = append(texts, offenceText, defenceText)
	if isPowerGreater(offenceTechnique.Power, defenceTechnique.Power) {
		t.offenceTeam.IncScore()
		texts = append(texts, "角間「決まったぁぁーー! "+offenceCharacter.Name+"のシュートが炸裂!!」")
	} else {
		texts = append(texts, "角間「キーパーの"+defenceCharacter.Name+"がしっかりキャッチ!」")
	}
	texts = append(texts, "------------------------")
	return texts
}
