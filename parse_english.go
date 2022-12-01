package nameconv

import "strings"

func ParseEnglish(n string) Name {
	words := strings.Split(n, " ")
	for i := range words {
		filteredWord := filter(words[i], lowercaseLetters+capitalLetters)
		if len(filteredWord) > 0 {
			words[i] = strings.ToLower(filteredWord)
		}
	}
	return Name(words)
}
