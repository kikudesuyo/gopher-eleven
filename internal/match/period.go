package match

import "strconv"

type Period struct {
	turn  Turn
	state PeriodState
}

type PeriodState string

const (
	firstHalf  PeriodState = "FIRST_HALF"
	secondHalf PeriodState = "SECOND_HALF"
	overTime   PeriodState = "OVER_TIME"
)

func (p *Period) proceed(texts []string) ([]string, bool) {
	texts = append(texts, "第"+strconv.Itoa(p.turn.count)+"ターン")
	texts = p.turn.performTechnique(texts)
	isEnd := p.turn.isEnd()
	p.turn.swapOffenceAndDefence()
	p.turn.incCount()
	return texts, isEnd

}
