package logger

import (
	"bytes"
	"testing"
)

func TestError(t *testing.T) {
	var buff bytes.Buffer
	l := NewFromWriters(nil, &buff)
	l.Error("test message %s", "hello")
	const expected = "-----> ERROR: test message hello\n"
	actual := buff.String()
	if actual != expected {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected )
	}
}

func TestDebug(t *testing.T) {
	var buff bytes.Buffer
	l := NewFromWriters(&buff, nil )
	l.Debug("test message %s", "hello")
	expected := "-----> DEBUG: test message hello\n"
	actual := buff.String()
	if actual != expected {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}

	l = NewFromWriters(nil, nil )
	buff.Reset()
	l.Debug("test message %s", "hello")
	expected = ""
	actual = buff.String()
	if actual != expected {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
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

func TestWarning(t *testing.T) {
	var buff bytes.Buffer
	l := NewFromWriters(nil, &buff)
	l.Warning("test %s", "foo")
	expected := "-----> WARN: test foo\n"
	actual := buff.String()
	if expected != actual {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}
}

func TestInfo(t *testing.T)  {
	var buff bytes.Buffer
	l := NewFromWriters(nil, &buff)
	l.Info("test %s", "bar")
	expected := "-----> INFO: test bar\n"
	actual := buff.String()
	if expected != actual {
		t.Fatalf("Actual: %q\nExpected: %q", actual, expected)
	}
}

type mockIndentity struct {}
func(mockIndentity) Identity()(string, string) {
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
