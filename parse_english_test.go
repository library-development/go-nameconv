package nameconv_test

import (
	"testing"

	"github.com/library-development/go-english"
	"github.com/library-development/go-nameconv"
)

func TestParseEnglish(t *testing.T) {
	n, err := nameconv.ParseEnglish("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParseEnglish() error = %v, want %v", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(english.Name{}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, english.Name{})
	}

	n, err = nameconv.ParseEnglish("hello")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello"})) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.ConvertStringListToName([]string{"hello"}))
	}

	n, err = nameconv.ParseEnglish("hello world")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello", "world"})) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.ConvertStringListToName([]string{"hello", "world"}))
	}

	n, err = nameconv.ParseEnglish("hello, world!")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello", "world"})) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.ConvertStringListToName([]string{"hello", "world"}))
	}

	n, err = nameconv.ParseEnglish("hello, world! 123")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.ConvertStringListToName([]string{"hello", "world", "123"})) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.ConvertStringListToName([]string{"hello", "world", "123"}))
	}
}
