package nameconv

import (
	"lib.dev/english"
)

func ConvertStringListToName(list []string) english.Name {
	n := english.Name{}
	for i := range list {
		n.AppendWord(list[i])
	}
	return n
}
