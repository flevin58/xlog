package xlog

import (
	"os"
	"strings"
	"testing"
)

func captureStderr(f func()) string {
	SetFlags(0)
	buf := new(strings.Builder)
	SetOutput(buf)
	f()
	SetOutput(os.Stderr)
	return strings.TrimSpace(buf.String())
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestDefaultPrefixes(t *testing.T) {
	expected := "ERROR: testing error string"
	result := captureStderr(func() {
		Error("testing error string")
	})
	if result != expected {
		Fatalf("expected: <%s>, got: <%s>\n", expected, result)
	}
}
func TestShortPrefixes(t *testing.T) {
	SetShortPrefixes()
	expected := "W: testing warning string"
	result := captureStderr(func() {
		Warn("testing warning string")
	})
	if result != expected {
		Fatalf("expected: <%s>, got: <%s>\n", expected, result)
	}
}

func TestPanic(t *testing.T) {
	assertPanic(t, func() {
		Panicln("This call panicked")
	})
}
