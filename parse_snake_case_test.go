package nameconv_test

import (
	"testing"

	"lib.dev/nameconv"
)

func TestParseSnakeCase(t *testing.T) {
	n, err := nameconv.ParseSnakeCase("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParseSnakeCase() error = %v, want %v", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(nil) {
		t.Errorf("ParseSnakeCase() = %v, want %v", n, nil)
	}

	n, err = nameconv.ParseSnakeCase("hello")
	if err != nil {
		t.Errorf("ParseSnakeCase() error = %v", err)
	}
	if !n.Equal(&nameconv.Name{[]string{"hello"}}) {
		t.Errorf("ParseSnakeCase() = %v, want %v", n, &nameconv.Name{[]string{"hello"}})
	}

	n, err = nameconv.ParseSnakeCase("hello_world")
	if err != nil {
		t.Errorf("ParseSnakeCase() error = %v", err)
	}
	if !n.Equal(&nameconv.Name{[]string{"hello", "world"}}) {
		t.Errorf("ParseSnakeCase() = %v, want %v", n, &nameconv.Name{[]string{"hello", "world"}})
	}
}
