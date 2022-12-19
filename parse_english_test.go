package nameconv_test

import (
	"testing"

	"github.com/library-development/go-nameconv"
)

func TestParseEnglish(t *testing.T) {
	n, err := nameconv.ParseEnglish("")
	if err != nameconv.ErrEmptyString {
		t.Errorf("ParseEnglish() error = %v, want %v", err, nameconv.ErrEmptyString)
	}
	if !n.Equal(nameconv.Name{}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.Name{})
	}

	n, err = nameconv.ParseEnglish("hello")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.Name{"hello"}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.Name{"hello"})
	}

	n, err = nameconv.ParseEnglish("hello world")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.Name{"hello", "world"}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.Name{"hello", "world"})
	}

	n, err = nameconv.ParseEnglish("hello, world!")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.Name{"hello", "world"}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.Name{"hello", "world"})
	}

	n, err = nameconv.ParseEnglish("hello, world! 123")
	if err != nil {
		t.Errorf("ParseEnglish() error = %v", err)
	}
	if !n.Equal(nameconv.Name{"hello", "world", "123"}) {
		t.Errorf("ParseEnglish() = %v, want %v", n, nameconv.Name{"hello", "world", "123"})
	}
}
