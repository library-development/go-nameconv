package nameconv

import "strings"

func ParsePascalCase(n string) Name {
	words := []string{}
	for i := 0; i < len(n); i++ {
		if n[i] >= 'A' && n[i] <= 'Z' {
			words = append(words, strings.ToLower(string(n[i])))
		} else {
			words[len(words)-1] += string(n[i])
		}
	}
	return Name(words)
}
