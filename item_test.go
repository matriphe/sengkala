package sengkala

import (
	"fmt"
	"sengkala/data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	RandomizerMock struct {
		mock.Mock
	}

	expected struct {
		expectedYear     int
		expectedSengkala string
		expectedMeaning  map[string]string
	}
)

var testItemData = map[int]expected{
	1984: {
		expectedYear:     1984,
		expectedSengkala: "Udaka Tekèk Wilasita Pamasé",
		expectedMeaning: map[string]string{
			"Udaka":    "air",
			"Tekèk":    "tokek",
			"Wilasita": "liang, liang kumbang",
			"Pamasé":   "raja",
		},
	},
	2020: {
		expectedYear:     2020,
		expectedSengkala: "Gegana Sikara Gegana Sikara",
		expectedMeaning: map[string]string{
			"Gegana": "angkasa, langit",
			"Sikara": "pengacauan, tangan, campur tangan.",
		},
	},
}

func (m *RandomizerMock) GetRandom(maxNum int) int {
	return 2
}

func TestItem(t *testing.T) {
	d := dictionary{
		watak:      data.Watak,
		meaning:    data.Meaning,
		Randomizer: new(RandomizerMock),
	}

	for year, expected := range testItemData {
		t.Run(fmt.Sprintf("item_%d", year), func(t *testing.T) {
			item := newItem(year, d)

			assert.Equal(t, expected.expectedYear, item.GetYear())
			assert.Equal(t, expected.expectedSengkala, item.GetSengkala())
			assert.Equal(t, expected.expectedMeaning, item.GetMeaning())
		})
	}
}
