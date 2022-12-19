package nameconv

import (
	"strings"
	"unicode"
)

// ValidateKebabCase returns an error if the string is not in kebab case.
func ValidateKebabCase(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	if s[0] == '-' || s[len(s)-1] == '-' {
		return ErrExpectedKebabCase
	}
	if strings.Contains(s, "--") {
		return ErrExpectedKebabCase
	}
	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) && r != '-' {
			return ErrExpectedKebabCase
		}
	}
	return nil
}
