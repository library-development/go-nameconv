package nameconv

import "strings"

// ParseEnglish parses names that use the English alphabet and spaces to separate words.
// If there is a non-letter character in the name, it will be ignored.
// If the name is empty, an error is returned.
func ParseEnglish(n string) (Name, error) {
	if len(n) == 0 {
		return Name{}, ErrEmptyString
	}
	words := strings.Split(n, " ")
	for i := range words {
		filteredWord := filter(words[i], lowercaseLetters+capitalLetters)
		if len(filteredWord) > 0 {
			words[i] = strings.ToLower(filteredWord)
		}
	}
	return Name{words}, nil
}
