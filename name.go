package nameconv

import (
	"strings"
	"unicode"
)

type Name []string

// AppendRune appends a rune to the end of the last word in the name.
func (n Name) AppendRune(r rune) {
	r = unicode.ToLower(r)
	if len(n) == 0 {
		n = Name{string(r)}
	}
	n[len(n)-1] += string(r)
}

// AppendWord appends a word to the end of the name.
func (n Name) AppendWord(s string) {
	n = append(n, s)
}

func (n Name) Equal(other Name) bool {
	if len(n) != len(other) {
		return false
	}
	for i := range n {
		if n[i] != other[i] {
			return false
		}
	}
	return true
}

func (n Name) String() string {
	return strings.Join(n, " ")
}

func (n Name) PascalCase() string {
	for i := range n {
		n[i] = strings.Title(n[i])
	}
	return strings.Join(n, "")
}

func (n Name) CamelCase() string {
	for i := range n {
		if i == 0 {
			n[i] = strings.ToLower(n[i])
		} else {
			n[i] = strings.Title(n[i])
		}
	}
	return strings.Join(n, "")
}

func (n Name) SnakeCase() string {
	for i := range n {
		n[i] = strings.ToLower(n[i])
	}
	return strings.Join(n, "_")
}

func (n Name) KebabCase() string {
	for i := range n {
		n[i] = strings.ToLower(n[i])
	}
	return strings.Join(n, "-")
}

func (n Name) AllLowerNoSpaces() string {
	return strings.Join(n, "")
}
