package match

import (
	"fmt"
	"strconv"

	"github.com/kikudesuyo/gopher-eleven/internal/display"
	"github.com/kikudesuyo/gopher-eleven/internal/team"
)

type Period string

var (
	firstHalf  Period = "FIRST_HALF"
	secondHalf Period = "SECOND_HALF"
	overTime   Period = "OVER_TIME"
	pk         Period = "PK"
	matchEnd   Period = "MATCH_END"
)

func (p *Period) setNext() {
	switch *p {
	case firstHalf:
		*p = secondHalf
	case secondHalf:
		*p = matchEnd //現段階では後半戦まで実装
	case overTime:
		*p = pk
	case pk:
		*p = matchEnd
	default:
		fmt.Println("invalid period state.")
	}
}

func (p *Period) isPeriodEnd(turnCount int) bool {
	if (*p == firstHalf) && (turnCount == 4) {
		return true
	}
	if (*p == secondHalf) && (turnCount == 4) {
		return true
	}
	return false
}

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

func (m *Match) appendEndFirstHalfPeriodTexts(texts []string) []string {
	texts = append(texts, "ホイッスル < ピッピーーーー", "角馬「ここで前半終了のホイッスルーーー!")
	scoreDiff := m.playerTeam.Score - m.opponentTeam.Score
	if scoreDiff > 0 {
		texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(scoreDiff)+"点のリードです!」")
	} else if scoreDiff == 0 {
		texts = append(texts, "\t"+m.playerTeam.Name+" 同点での折り返しです!」")
	} else {
		texts = append(texts, "\t"+m.playerTeam.Name+" "+strconv.Itoa(-1*scoreDiff)+"点のビハインドです..!」")
	}
	texts = append(texts, "===========================================")
	return texts
}

func (m *Match) appendEndMatchTexts(texts []string) []string {
	if m.playerTeam.Score > m.opponentTeam.Score {
		texts = append(texts, "ホイッスル < ピッピーーーー", "角馬「ここ試合終了のホイッスルーーー!")
		texts = append(texts, "\t"+strconv.Itoa(m.playerTeam.Score)+"対"+strconv.Itoa(m.opponentTeam.Score)+"で"+m.playerTeam.Name+"の勝利です!!!」")
	} else if m.playerTeam.Score == m.opponentTeam.Score {
		texts = append(texts, "\t"+m.playerTeam.Name+" 同点で延長戦に突入!!", "\t              っっっああああっと!"+m.opponentTeam.Name+"が試合を放棄!!!", "\t       よって"+m.playerTeam.Name+"の勝利です!!!」")
	} else {
		texts = append(texts, "ホイッスル < ピッピーーーー", "角馬「ここ試合終了のホイッスルーーー!")
		texts = append(texts, "\t"+strconv.Itoa(m.opponentTeam.Score)+"対"+strconv.Itoa(m.playerTeam.Score)+"で"+m.opponentTeam.Name+"の勝利です..."+"\t惜しくも"+m.playerTeam.Name+"は敗退です...!」")
	}
	return texts
}

func (m *Match) Proceed() (display.Display, bool) {
	var texts []string
	if m.turn.count == 1 && m.period == secondHalf {
		texts = append(texts, "ホイッスル < ピーーーーーー", "角馬「後半戦スタートです!」")
	}
	texts, turnCount := m.turn.proceed(texts)
	if !m.period.isPeriodEnd(turnCount) {
		disp := display.NewDisplay(texts...)
		return disp, false
	}

	if m.period == firstHalf {
		m.turn.resetCount()
		texts = m.appendEndFirstHalfPeriodTexts(texts)
		m.period.setNext()
	} else if m.period == secondHalf {
		m.turn.resetCount()
		texts = m.appendEndMatchTexts(texts)
		m.period.setNext()
	}
	disp := display.NewDisplay(texts...)
	return disp, m.period == matchEnd
}
