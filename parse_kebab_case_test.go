package nameconv_test

import (
	"testing"

	"lib.dev/nameconv"
)

func TestParseKebabCase(t *testing.T) {
	n, err := nameconv.ParseKebabCase("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParseKebabCase() error = \"%v\", want \"%v\"", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(nil) {
		t.Errorf("ParseKebabCase() = \"%v\", want \"%v\"", n, nil)
	}

	n, err = nameconv.ParseKebabCase("hello")
	if err != nil {
		t.Errorf("ParseKebabCase() error = \"%v\"", err)
	}
	if !n.Equal(&nameconv.Name{[]string{"hello"}}) {
		t.Errorf("ParseKebabCase() = \"%v\", want \"%v\"", n, &nameconv.Name{[]string{"hello"}})
	}

	n, err = nameconv.ParseKebabCase("hello-world")
	if err != nil {
		t.Errorf("ParseKebabCase() error = \"%v\"", err)
	}
	if !n.Equal(&nameconv.Name{[]string{"hello", "world"}}) {
		t.Errorf("ParseKebabCase() = \"%v\", want \"%v\"", n, &nameconv.Name{[]string{"hello", "world"}})
	}
}
