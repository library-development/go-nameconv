package nameconv

import "strings"

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) ParseEnglish() []string {
	words := strings.Split(n.String(), " ")
	for i := range words {
		filteredWord := filter(words[i], lowercaseLetters+capitalLetters)
		if len(filteredWord) > 0 {
			words[i] = strings.ToLower(filteredWord)
		}
	}
	return words
}

func (n Name) PascalCase() string {
	words := n.ParseEnglish()
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func (n Name) CamelCase() string {
	words := n.ParseEnglish()
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

func (n Name) SnakeCase() string {
	words := n.ParseEnglish()
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}

func (n Name) KebabCase() string {
	words := n.ParseEnglish()
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "-")
}

func (n Name) AllLowerNoSpaces() string {
	words := n.ParseEnglish()
	return strings.Join(words, "")
}
