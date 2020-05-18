package sengkala

import (
	"fmt"
	"sengkala/data"
	"strconv"
	"strings"
)

type item struct {
	year     int
	err      error
	sengkala []string
	meaning  map[string]map[string]string
}

// GetSengkala returns sengkala words
func GetSengkala(year string, randomizer Randomizer) []string {
	res := make([]string, len(year))
	for i, t := range reverse(year) {
		it, _ := strconv.Atoi(fmt.Sprintf("%c", t))
		w := data.Watak[it]
		ch := w[randomizer.GetRandom(len(w)-1)]
		res[i] = strings.Title(ch[randomizer.GetRandom(len(ch)-1)])
	}

	return res
}

func newItem(year string) *item {
	y, err := strconv.Atoi(year)
	if err != nil {
		return &item{
			err: err,
		}
	}

	return &item{
		year:     y,
		err:      nil,
		sengkala: GetSengkala(year, newRandomizer()),
		meaning:  nil,
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
