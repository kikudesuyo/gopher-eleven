package match

import "math/rand"

func isPowerGreater(basePower int, opponentPower int) bool {
	powerDiff := basePower - opponentPower
	return isGreater(powerDiff)
}

func isGreater(powerDiff int) bool {
	if powerDiff > 0 {
		return rand.Intn(10) > 3
	} else {
		return rand.Intn(10) > 7
	}
}
