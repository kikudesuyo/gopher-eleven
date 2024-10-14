package team

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Inc(t *testing.T) {
	team := Team{
		Characters: nil,
		Name:       "",
		Score:      2,
	}
	got := team.Inc()
	want := 3
	assert.Equal(t, want, got)
}
