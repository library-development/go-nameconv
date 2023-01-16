package nameconv

import (
	"strings"

	"lib.dev/english"
)

// ParseKebabCase parses a kebab-case string into a Name.
// It returns an error if the string is empty or not kebab-case.
func ParseKebabCase(s string) (english.Name, error) {
	err := validateKebabCase(s)
	if err != nil {
		return nil, err
	}
	return ConvertStringListToName(strings.Split(s, "-")), nil
}
