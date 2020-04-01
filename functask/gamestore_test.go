package functask_test

import (
	f "github.com/stevedesilva/gofuncPointers/functask"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := f.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
