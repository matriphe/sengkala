package sengkala

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom_GetRandom(t *testing.T) {
	randomizer := newRandomizer()
	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			actual := randomizer.GetRandom(3)
			assert.True(t, 0 <= actual && actual <= 3)
		})
	}
}
