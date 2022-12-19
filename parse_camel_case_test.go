package nameconv

import "testing"

func TestParseCamelCase(t *testing.T) {
	n, err := ParseCamelCase("")
	if err != ErrEmptyString {
		t.Errorf("ParseCamelCase() error = \"%v\", want \"%v\"", err, ErrEmptyString)
	}
	if !n.Equal(nil) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, nil)
	}

	n, err = ParseCamelCase("hello")
	if err != nil {
		t.Errorf("ParseCamelCase() error = \"%v\"", err)
	}
	if !n.Equal(&Name{[]string{"hello"}}) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, &Name{[]string{"hello"}})
	}

	n, err = ParseCamelCase("helloWorld")
	if err != nil {
		t.Errorf("ParseCamelCase() error = \"%v\"", err)
	}
	if !n.Equal(&Name{[]string{"hello", "world"}}) {
		t.Errorf("ParseCamelCase() = \"%v\", want \"%v\"", n, &Name{[]string{"hello", "world"}})
	}
}
