package nameconv

import "strings"

func ParseSnakeCase(n string) Name {
	words := strings.Split(n, "_")
	return Name(words)
}
