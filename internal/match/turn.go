package match

import (
	"strconv"

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

func (t *Turn) resetCount() {
	t.count = 1
}

func (t *Turn) swapOffenceAndDefence() {
	t.offenceTeam, t.defenceTeam = t.defenceTeam, t.offenceTeam
}

func (t *Turn) performTechnique(texts []string) []string {
	offenceCharacter := t.offenceTeam.Characters[1]
	offenceTechnique, offenceText := offenceCharacter.Perform()
	defenceCharacter := t.defenceTeam.Characters[0]
	defenceTechnique, defenceText := defenceCharacter.Perform()
	texts = append(texts, "第"+strconv.Itoa(t.count)+"ターン")
	texts = append(texts, "角馬「"+t.offenceTeam.Name+"の"+offenceCharacter.Name+"のシュートだ!!」")
	texts = append(texts, offenceText, defenceText)
	if isPowerGreater(offenceTechnique.Power, defenceTechnique.Power) {
		t.offenceTeam.IncScore()
		texts = append(texts, "角馬「決まったぁぁーー! "+offenceCharacter.Name+"のシュートが炸裂!!」")
	} else {
		texts = append(texts, "角馬「キーパーの"+defenceCharacter.Name+"がしっかりキャッチ!」")
	}
	texts = append(texts, "------------------------")
	return texts
}

func (t *Turn) proceed(texts []string) ([]string, int) {
	texts = t.performTechnique(texts)
	t.swapOffenceAndDefence()
	t.incCount()
	return texts, t.count
}
