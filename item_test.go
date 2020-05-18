package sengkala

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RandomizerMock struct {
	mock.Mock
}

func (m *RandomizerMock) GetRandom(maxNum int) int {
	return 2
}

func TestGetSengkala(t *testing.T) {
	for year, expected := range map[string][]string{
		"1984": {"Udaka", "Tekèk", "Wilasita", "Pamasé"},
		"2020": {"Gegana", "Sikara", "Gegana", "Sikara"},
	} {
		t.Run(fmt.Sprintf("year_%s", year), func(t *testing.T) {
			actual := GetSengkala(year, new(RandomizerMock))
			assert.Equal(t, expected, actual)
		})
	}
}

func TestItem(t *testing.T) {
	_, abcErr := strconv.Atoi("abc")
	_, emptyErr := strconv.Atoi("")

	for year, expected := range map[string]item{
		"1984": {
			year:     1984,
			err:      nil,
		},
		"abc": {
			err: abcErr,
		},
		"": {
			err: emptyErr,
		},
	} {
		t.Run(fmt.Sprintf("year_%s", year), func(t *testing.T) {
			actual := newItem(year)
			assert.Equal(t, expected.year, actual.year)
			assert.Equal(t, expected.err, actual.err)
		})
	}
}
