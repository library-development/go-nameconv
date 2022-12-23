package nameconv

import (
	"strings"
	"unicode"
)

// ParseCamelCase parses a camelCase name into a Name.
// If the name is not valid camelCase, an error is returned.
func ParseCamelCase(s string) (*Name, error) {
	err := ValidateCamelCase(s)
	if err != nil {
		return nil, err
	}
	n := &Name{
		Words: []string{},
	}
	for _, r := range s {
		if unicode.IsUpper(r) {
			st := strings.ToLower(string(r))
			n.AppendWord(st)
		} else {
			n.AppendRune(r)
		}
	}
	return n, nil
}
