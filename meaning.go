package sengkala

import "strings"

// GetMeaning returns meaning from dictionary
func GetMeaning(word string, dictionary map[string]string) string {
	word = strings.ToLower(word)

	if mean, ok := dictionary[word]; ok {
		return mean
	}

	return ""
}
