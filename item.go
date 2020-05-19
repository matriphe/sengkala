package sengkala

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matriphe/sengkala/data"
)

type (
	item struct {
		year     int
		sengkala []string
		meaning  map[string]string
	}

	dictionary struct {
		watak      data.WatakType
		meaning    data.MeaningType
		Randomizer Randomizer
	}

	// Item interfaces item
	Item interface {
		GetYear() int
		GetSengkala() string
		GetMeaning() map[string]string
	}
)

func newItem(year int, dict dictionary) *item {
	sengkala := getSengkala(year, dict.watak, dict.Randomizer)
	meanings := getMeaning(sengkala, dict.meaning)

	return &item{
		year:     year,
		sengkala: sengkala,
		meaning:  meanings,
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func getSengkala(y int, watak data.WatakType, randomizer Randomizer) []string {
	year := strconv.Itoa(y)
	res := make([]string, len(year))
	for i, t := range reverse(year) {
		it, _ := strconv.Atoi(fmt.Sprintf("%c", t))
		w := watak[it]
		ch := w[randomizer.GetRandom(len(w)-1)]
		res[i] = strings.Title(ch[randomizer.GetRandom(len(ch)-1)])
	}

	return res
}

func getMeaning(words []string, meaning data.MeaningType) map[string]string {
	res := make(map[string]string, len(words))
	for _, word := range words {
		if v, ok := meaning[strings.ToLower(word)]["id"]; ok {
			res[word] = v
		} else {
			res[word] = ""
		}
	}

	return res
}

// GetYear returns year
func (i *item) GetYear() int {
	return i.year
}

// GetSengkala returns sengkala
func (i *item) GetSengkala() string {
	return strings.Join(i.sengkala, " ")
}

// GetMeaning returns sengkala meaning
func (i *item) GetMeaning() map[string]string {
	return i.meaning
}
