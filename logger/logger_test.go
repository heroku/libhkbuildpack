package logger

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevelFormat(t *testing.T) {
	var stderr, stdout bytes.Buffer
	l := NewFromWriters(&stderr, &stdout)
	tt := []struct {
		name     string
		testFn   func()
		out      fmt.Stringer
		expected string
	}{
		{
			name: "info",
			testFn: func() {
				l.Info("test %s", "foo")
			},
			out:      &stdout,
			expected: "-----> INFO:  test foo\n",
		},
		{
			name: "info empty",
			testFn: func() {
				l.Info("")
			},
			out:      &stdout,
			expected: "\n",
		},
		{
			name: "warn",
			testFn: func() {
				l.Warning("test %s", "warning")
			},
			out:      &stdout,
			expected: "-----> WARN:  test warning\n",
		},
		{
			name: "warn empty",
			testFn: func() {
				l.Warning("")
			},
			out:      &stdout,
			expected: "\n",
		},
		{
			name: "error",
			testFn: func() {
				l.Error("test %s", "error")
			},
			out:      &stdout,
			expected: "-----> ERROR: test error\n",
		},
		{
			name: "error empty",
			testFn: func() {
				l.Error("")
			},
			out:      &stdout,
			expected: "\n",
		},
		{
			name: "debug",
			testFn: func() {
				l.Debug("test %s", "debug")
			},
			out:      &stderr,
			expected: "-----> DEBUG: test debug\n",
		},
		{
			name: "debug empty",
			testFn: func() {
				l.Debug("")
			},
			out:      &stderr,
			expected: "\n",
		},
	}
	for _, tc := range tt {
		stderr.Reset()
		stdout.Reset()
		t.Run(tc.name, func(t *testing.T) {
			tc.testFn()
			assert(t, tc.out.String(), tc.expected)

		})
	}
}

func assert(t *testing.T, actual, expected string) {
	if expected != actual {
		t.Logf("expected %q", expected)
		t.Logf("actual   %q", actual)
		t.Fatal()
	}
}

func TestFirstLine(t *testing.T) {
	var buff bytes.Buffer
	l := NewFromWriters(nil, &buff)
	l.FirstLine("test %s", "hello")
	expected := "-----> test hello\n"
	actual := buff.String()
	if expected != actual {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}
}

func TestSubsequentLine(t *testing.T) {
	var buff bytes.Buffer
	l := NewFromWriters(nil, &buff)
	l.SubsequentLine("test %s", "foo")
	expected := "      test foo\n"
	actual := buff.String()
	if actual != expected {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}
}

type mockIndentity struct{}

func (mockIndentity) Identity() (string, string) {
	return "foo", "bar"
}

func TestPrettyIndentity(t *testing.T) {
	var m mockIndentity
	l := NewFromWriters(nil, nil)
	expected := "foo - bar"
	actual := l.PrettyIdentity(m)
	if expected != actual {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}
}
