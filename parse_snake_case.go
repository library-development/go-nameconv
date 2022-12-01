package nameconv

import "strings"

func ParseSnakeCase(n string) []string {
	words := strings.Split(n, "_")
	return Name(words)
}
