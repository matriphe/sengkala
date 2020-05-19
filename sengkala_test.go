package sengkala

import (
	"fmt"
	"sengkala/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSengkalaData = map[string]struct {
	changedYear string
	old         struct {
		expectedSurya  expected
		expectedCandra expected
	}
	new struct {
		expectedSurya  expected
		expectedCandra expected
	}
}{
	"1983": {
		changedYear: "1984",
		old: struct {
			expectedSurya  expected
			expectedCandra expected
		}{
			expectedSurya: expected{
				expectedYear:     1983,
				expectedSengkala: "Bentèr Tekèk Wilasita Pamasé",
				expectedMeaning: map[string]string{
					"Bentèr":   "panas",
					"Pamasé":   "raja",
					"Tekèk":    "tokek",
					"Wilasita": "liang, liang kumbang",
				},
			},
			expectedCandra: expected{
				expectedYear:     1915,
				expectedSengkala: "Sara Pamasé Wilasita Pamasé",
				expectedMeaning: map[string]string{
					"Pamasé":   "raja",
					"Sara":     "senjata, panah",
					"Wilasita": "liang, liang kumbang",
				},
			},
		},
		new: struct {
			expectedSurya  expected
			expectedCandra expected
		}{
			expectedSurya: expected{
				expectedYear:     1984,
				expectedSengkala: "Udaka Tekèk Wilasita Pamasé",
				expectedMeaning: map[string]string{
					"Udaka":    "air",
					"Tekèk":    "tokek",
					"Wilasita": "liang, liang kumbang",
					"Pamasé":   "raja",
				},
			},
			expectedCandra: expected{
				expectedYear:     1916,
				expectedSengkala: "Karaséng Pamasé Wilasita Pamasé",
				expectedMeaning: map[string]string{
					"Karaséng": "terasa oleh ( terasa pada )",
					"Pamasé":   "raja",
					"Wilasita": "liang, liang kumbang",
				},
			},
		},
	},
}

func TestSengkala(t *testing.T) {
	d := dictionary{
		watak:      data.Watak,
		meaning:    data.Meaning,
		Randomizer: new(RandomizerMock),
	}

	for year, expected := range testSengkalaData {
		t.Run(fmt.Sprintf("sengkala_%s", year), func(t *testing.T) {
			sengkala := FromYear(year)
			sengkala.SetDictionary(d)

			assert.NoError(t, sengkala.GetError())

			assert.Equal(t, expected.old.expectedSurya.expectedYear, sengkala.GetSuryaSengkala().GetYear())
			assert.Equal(t, expected.old.expectedSurya.expectedSengkala, sengkala.GetSuryaSengkala().GetSengkala())
			assert.Equal(t, expected.old.expectedSurya.expectedMeaning, sengkala.GetSuryaSengkala().GetMeaning())

			assert.Equal(t, expected.old.expectedCandra.expectedYear, sengkala.GetCandraSengkala().GetYear())
			assert.Equal(t, expected.old.expectedCandra.expectedSengkala, sengkala.GetCandraSengkala().GetSengkala())
			assert.Equal(t, expected.old.expectedCandra.expectedMeaning, sengkala.GetCandraSengkala().GetMeaning())

			sengkala.SetYear(expected.changedYear)

			assert.Equal(t, expected.new.expectedSurya.expectedYear, sengkala.GetSuryaSengkala().GetYear())
			assert.Equal(t, expected.new.expectedSurya.expectedSengkala, sengkala.GetSuryaSengkala().GetSengkala())
			assert.Equal(t, expected.new.expectedSurya.expectedMeaning, sengkala.GetSuryaSengkala().GetMeaning())

			assert.Equal(t, expected.new.expectedCandra.expectedYear, sengkala.GetCandraSengkala().GetYear())
			assert.Equal(t, expected.new.expectedCandra.expectedSengkala, sengkala.GetCandraSengkala().GetSengkala())
			assert.Equal(t, expected.new.expectedCandra.expectedMeaning, sengkala.GetCandraSengkala().GetMeaning())
		})
	}
}
