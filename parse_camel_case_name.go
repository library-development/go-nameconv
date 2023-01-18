package nameconv

import (
	"strings"
	"unicode"

	"github.com/library-development/go-english"
)

// ParseCamelCase parses a camelCase name into a Name.
// If the name is not valid camelCase, an error is returned.
func ParseCamelCase(s string) (english.Name, error) {
	err := validateCamelCase(s)
	if err != nil {
		return nil, err
	}
	n := english.Name{}
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
