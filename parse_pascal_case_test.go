package nameconv_test

import (
	"testing"

	"lib.dev/nameconv"
)

func TestParsePascalCase(t *testing.T) {
	n, err := nameconv.ParsePascalCase("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParsePascalCase() error = \"%v\", want \"%v\"", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(nil) {
		t.Errorf("ParsePascalCase() = \"%v\", want \"%v\"", n, nil)
	}

	n, err = nameconv.ParsePascalCase("hello")
	if err != nameconv.ErrExpectedPascalCase {
		t.Errorf("ParsePascalCase() error = \"%v\", want \"%v\"", err, nameconv.ErrExpectedPascalCase)
	}
	if !n.Equal(nil) {
		t.Errorf("ParsePascalCase() = \"%v\", want \"%v\"", n, nil)
	}

	n, err = nameconv.ParsePascalCase("Hello")
	if err != nil {
		t.Errorf("ParsePascalCase() error = \"%v\"", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello"})) {
		t.Errorf("ParsePascalCase() = \"%v\", want \"%v\"", n, nameconv.ConvertStringListToName([]string{"hello"}))
	}

	n, err = nameconv.ParsePascalCase("HelloWorld")
	if err != nil {
		t.Errorf("ParsePascalCase() error = \"%v\"", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello", "world"})) {
		t.Errorf("ParsePascalCase() = \"%v\", want \"%v\"", n, nameconv.ConvertStringListToName([]string{"hello", "world"}))
	}
}
