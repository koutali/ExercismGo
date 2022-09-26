package isogram

import "strings"

func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	occurences := make(map[rune]int)
	for _, char := range word {
		if char == '-' || char == ' ' {
			continue
		}

		occurences[char]++
	}

	for _, value := range occurences {
		if value > 1 {
			return false
		}
	}
	return true
}
