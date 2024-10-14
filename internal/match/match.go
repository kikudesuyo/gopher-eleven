package match

import (
	"strconv"

	"github.com/kikudesuyo/gopher-eleven/internal/character"
	"github.com/kikudesuyo/gopher-eleven/internal/display"
	"github.com/kikudesuyo/gopher-eleven/internal/team"
)

type Turn struct {
	count       int
	offenceTeam *team.Team
	defenceTeam *team.Team
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
	playerTeam   *team.Team
	opponentTeam *team.Team
}

func InitMatch() Match {
	playerTeam := team.GetPlayerTeam()
	opponentTeam := team.GetOpponentTeam()
	match := Match{
		turn:         Turn{1, &playerTeam, &opponentTeam},
		period:       firstHalf,
		playerTeam:   &playerTeam,
		opponentTeam: &opponentTeam,
	}
	return match
}

func (m *Match) performTechnique(texts []string) (character.Character, character.Technique, character.Character, character.Technique, []string) {
	offenceCharacter := m.turn.offenceTeam.Characters[1]
	offenceTechnique := offenceCharacter.Perform()
	defenceCharacter := m.turn.defenceTeam.Characters[0]
	defenceTechnique := defenceCharacter.Perform()
	texts = append(texts, offenceCharacter.Name+"「"+offenceTechnique.Name+"!!」",
		defenceCharacter.Name+"「"+defenceTechnique.Name+"!!」",
	)
	return offenceCharacter, offenceTechnique, defenceCharacter, defenceTechnique, texts
}

func (m *Match) appendTurnResultTexts(texts []string, offencePower, defencePower int, offenceCharacterName, defenceCharacterName string) []string {
	if isPowerGreater(offencePower, defencePower) {
		m.turn.offenceTeam.Inc()
		texts = append(texts, "角間「決まったぁぁーー! "+offenceCharacterName+"のシュートが炸裂!!」")
	} else {
		texts = append(texts, "角間「キーパーの"+defenceCharacterName+"がしっかりキャッチ!」")
	}
	texts = append(texts, "------------------------")
	return texts
}

func (m *Match) appendEndPeriodText(texts []string) []string {
	texts = append(texts, "ホイッスル < ピッピーーーー", "角間「ここで前半終了のホイッスルーーー!")
	scoreDiff := m.playerTeam.Score - m.opponentTeam.Score
	if scoreDiff > 0 {
		texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(scoreDiff)+"点のリードです!」")
	} else if scoreDiff == 0 {
		texts = append(texts, "\t"+m.playerTeam.Name+" 同点での折り返しです!」")
	} else {
		texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(-1*scoreDiff)+"点のビハインドです..!」")
	}
	return texts
}

func (m *Match) Proceed() (display.Display, bool, error) {
	var texts []string
	texts = append(texts, "第"+strconv.Itoa(m.turn.count)+"ターン")
	offenceCharacter, offenceTechnique, defenceCharacter, defenceTechnique, texts := m.performTechnique(texts)
	texts = m.appendTurnResultTexts(
		texts,
		offenceTechnique.Power,
		defenceTechnique.Power,
		offenceCharacter.Name,
		defenceCharacter.Name,
	)
	isEnd := m.turn.count == 4
	if isEnd {
		texts = m.appendEndPeriodText(texts)
	}
	m.turn.offenceTeam, m.turn.defenceTeam = m.turn.defenceTeam, m.turn.offenceTeam
	m.turn.Inc()

	disp := display.NewDisplay(texts...)
	return disp, isEnd, nil
}
