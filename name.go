package nameconv

import (
	"strings"
	"unicode"
)

type Name struct {
	Words []string
}

// AppendRune appends a rune to the end of the last word in the name.
func (n *Name) AppendRune(r rune) {
	r = unicode.ToLower(r)
	if len(n.Words) == 0 {
		n.Words = []string{string(r)}
		return
	}
	n.Words[len(n.Words)-1] += string(r)
}

// AppendWord appends a word to the end of the name.
func (n *Name) AppendWord(s string) {
	n.Words = append(n.Words, s)
}

func (n *Name) Equal(other *Name) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n.Words) != len(other.Words) {
		return false
	}
	for i := range n.Words {
		if n.Words[i] != other.Words[i] {
			return false
		}
	}
	return true
}

func (n *Name) String() string {
	return strings.Join(n.Words, " ")
}

func (n *Name) PascalCase() string {
	for i := range n.Words {
		n.Words[i] = strings.Title(n.Words[i])
	}
	return strings.Join(n.Words, "")
}

func (n *Name) CamelCase() string {
	for i := range n.Words {
		if i == 0 {
			n.Words[i] = strings.ToLower(n.Words[i])
		} else {
			n.Words[i] = strings.Title(n.Words[i])
		}
	}
	return strings.Join(n.Words, "")
}

func (n *Name) SnakeCase() string {
	for i := range n.Words {
		n.Words[i] = strings.ToLower(n.Words[i])
	}
	return strings.Join(n.Words, "_")
}

func (n *Name) KebabCase() string {
	for i := range n.Words {
		n.Words[i] = strings.ToLower(n.Words[i])
	}
	return strings.Join(n.Words, "-")
}

func (n *Name) AllLowerNoSpaces() string {
	return strings.Join(n.Words, "")
}

func (n *Name) TitleCase() string {
	for i := range n.Words {
		n.Words[i] = strings.Title(n.Words[i])
	}
	return strings.Join(n.Words, " ")
}
