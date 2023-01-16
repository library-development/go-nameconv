package nameconv_test

import (
	"testing"

	"lib.dev/nameconv"
)

func TestParseCamelCase(t *testing.T) {
	n, err := nameconv.ParseCamelCase("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParseCamelCase() error = \"%v\", want \"%v\"", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(nil) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, nil)
	}

	n, err = nameconv.ParseCamelCase("hello")
	if err != nil {
		t.Errorf("ParseCamelCase() error = \"%v\"", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello"})) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, nameconv.ConvertStringListToName([]string{"hello"}))
	}

	n, err = nameconv.ParseCamelCase("helloWorld")
	if err != nil {
		t.Errorf("ParseCamelCase() error = \"%v\"", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello", "world"})) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, nameconv.ConvertStringListToName([]string{"hello", "world"}))
	}
}
