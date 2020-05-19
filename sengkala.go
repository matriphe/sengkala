package sengkala

import (
	"sengkala/data"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

const yearLayout = "2006"

type (
	// Sengkala holds sengkala data
	Sengkala struct {
		year  time.Time
		error error
		dict  dictionary
	}

	// SengkalaInterface interfaces sengkala
	SengkalaInterface interface {
		FromYear(year string) SengkalaInterface
		SetYear(year string)
		SetDictionary(dict dictionary)
		GetError() error
		GetSuryaSengkala() Item
		GetCandraSengkala() Item
	}
)

// FromYear creates Sengkala
func FromYear(year string) *Sengkala {
	y, err := time.Parse(yearLayout, year)
	if err != nil {
		return &Sengkala{error: err}
	}

	return &Sengkala{
		year:  y,
		error: nil,
		dict: dictionary{
			watak:      data.Watak,
			meaning:    data.Meaning,
			Randomizer: newRandomizer(),
		},
	}
}

// SetYear sets year
func (s *Sengkala) SetYear(year string) {
	y, err := time.Parse(yearLayout, year)
	if err != nil {
		s.error = err
	}

	s.year = y
}

// SetDictionary set dictionary
func (s *Sengkala) SetDictionary(dict dictionary) {
	s.dict = dict
}

// GetError gets error
func (s *Sengkala) GetError() error {
	return s.error
}

// GetSuryaSengkala returns surya sengkala
func (s *Sengkala) GetSuryaSengkala() Item {
	if s.error != nil {
		return nil
	}

	return newItem(s.year.Year(), s.dict)
}

// GetCandraSengkala returns candra sengkala
func (s *Sengkala) GetCandraSengkala() Item {
	if s.error != nil {
		return nil
	}

	hijriYear, _, _ := hijri.ToHijri(s.year)
	javaneseYear := hijriYear + 512

	return newItem(javaneseYear, s.dict)
}
