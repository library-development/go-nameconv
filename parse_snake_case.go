package nameconv

import "strings"

// ParseSnakeCase parses a snake_case string into a Name.
// If the string is not valid snake_case, an error is returned.
func ParseSnakeCase(s string) (*Name, error) {
	err := ValidateSnakeCase(s)
	if err != nil {
		return nil, err
	}
	return &Name{strings.Split(s, "_")}, nil
}
