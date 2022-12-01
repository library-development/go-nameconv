package nameconv

import "strings"

type Name []string

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
