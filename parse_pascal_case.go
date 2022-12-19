package nameconv

import (
	"strings"
	"unicode"
)

// ParsePascalCase parses a PascalCase name into a Name.
// If the name is not valid PascalCase, an error is returned.
func ParsePascalCase(s string) (*Name, error) {
	err := ValidatePascalCase(s)
	if err != nil {
		return nil, err
	}
	n := &Name{}
	for _, r := range s {
		if unicode.IsUpper(r) {
			s := strings.ToLower(string(r))
			n.AppendWord(s)
		} else {
			n.AppendRune(r)
		}
	}
	return n, nil
}
