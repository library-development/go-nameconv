package nameconv

import (
	"strings"

	"lib.dev/english"
)

// ParseSnakeCase parses a snake_case string into a Name.
// If the string is not valid snake_case, an error is returned.
func ParseSnakeCase(s string) (english.Name, error) {
	err := validateSnakeCase(s)
	if err != nil {
		return nil, err
	}
	return ConvertStringListToName(strings.Split(s, "_")), nil
}
