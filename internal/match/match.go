package match

import (
	"strconv"

	"github.com/kikudesuyo/gopher-eleven/internal/display"
	"github.com/kikudesuyo/gopher-eleven/internal/team"
)

type Match struct {
	period       Period
	playerTeam   *team.Team
	opponentTeam *team.Team
}

func InitMatch() Match {
	playerTeam := team.GetPlayerTeam()
	opponentTeam := team.GetOpponentTeam()
	match := Match{
		period:       Period{Turn{1, &playerTeam, &opponentTeam}, firstHalf},
		playerTeam:   &playerTeam,
		opponentTeam: &opponentTeam,
	}
	return match
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
	texts, isEnd := m.period.proceed(texts)
	if isEnd {
		texts = m.appendEndPeriodText(texts)
	}
	disp := display.NewDisplay(texts...)
	return disp, isEnd, nil
}
