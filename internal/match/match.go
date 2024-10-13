package match

import (
	"strconv"

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

func (m *Match) Proceed() (display.Display, bool, error) {
	offenceCharacter := m.turn.offenceTeam.Characters[1]
	offenceTechnique := offenceCharacter.Perform()
	defenceTeam := m.turn.defenceTeam
	defenceCharacter := defenceTeam.Characters[0]
	defenceTechnique := defenceCharacter.Perform()
	texts := []string{
		"第" + strconv.Itoa(m.turn.count) + "ターン",
		offenceCharacter.Name + "「" + offenceTechnique.Name + "!!」",
		defenceCharacter.Name + "「" + defenceTechnique.Name + "!!」",
	}
	if offenceTechnique.Power > defenceTechnique.Power {
		m.turn.offenceTeam.Inc()
		texts = append(texts, "角間「決まったー! "+offenceCharacter.Name+"のシュートが炸裂!!」")
	} else {
		texts = append(texts, "角間「キーパーの"+defenceCharacter.Name+"がしっかりキャッチ!」")
	}
	texts = append(texts, "------------------------")
	isEnd := m.turn.count == 3
	if isEnd {
		texts = append(texts, "ホイッスル < ピッピーーーー", "角間「ここで前半終了のホイッスルーーー!")
		scoreDiff := m.playerTeam.Score - m.opponentTeam.Score
		if scoreDiff > 0 {
			texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(scoreDiff)+"点のリードです!」")
		} else if scoreDiff == 0 {
			texts = append(texts, "\t"+m.playerTeam.Name+" 同点での折り返しです!」")
		} else {
			texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(scoreDiff)+"点のビハインドです..!」")
		}
	}
	disp := display.NewDisplay(texts...)
	m.turn.offenceTeam, m.turn.defenceTeam = m.turn.defenceTeam, m.turn.offenceTeam
	m.turn.Inc()
	return disp, isEnd, nil
}
