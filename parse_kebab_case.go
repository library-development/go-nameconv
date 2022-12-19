package nameconv

import "strings"

// ParseKebabCase parses a kebab-case string into a Name.
// It returns an error if the string is empty or not kebab-case.
func ParseKebabCase(s string) (*Name, error) {
	err := ValidateKebabCase(s)
	if err != nil {
		return nil, err
	}
	return &Name{strings.Split(s, "-")}, nil
}
