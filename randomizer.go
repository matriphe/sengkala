package sengkala

import (
	"math/rand"
	"time"
)

// Randomizer interface
type Randomizer interface {
	GetRandom(maxNum int) int
}

type random struct {
}

// newRandomizer return Randomizer concret (random)
func newRandomizer() *random {
	return &random{}
}

// GetRandom returns random number between 0 to maxNum
func (r *random) GetRandom(maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum)
}
