// Package logger provides console output for libhkbuildpack operations.
package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const (
	indent  = "      "
	prefix  = "----->"
	error   = "ERROR"
	warning = "WARN"
	debug   = "DEBUG"
	info    = "INFO"
)

// Log supports logging related methods and state
type Log struct {
	sync.Mutex
	info  *bufio.Writer
	debug *bufio.Writer
}

type logLevelEnalbler interface {
	IsDebugEnabled() bool
	IsInfoEnabled() bool
}

func New(l logLevelEnalbler) *Log {
	var logger Log
	if l == nil {
		return &logger
	}
	if l.IsDebugEnabled() {
		logger.debug = bufio.NewWriter(os.Stderr)
	}
	if l.IsInfoEnabled() {
		logger.info = bufio.NewWriter(os.Stdout)
	}
	return &logger
}

func NewFromWriters(debug, info io.Writer) *Log {
	var logger Log
	if info != nil {
		logger.info = bufio.NewWriter(info)
	}
	if debug != nil {
		logger.debug = bufio.NewWriter(debug)
	}
	return &logger
}

// Error prints an error message to the console if an info logger is provided.
func (l Log) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printInfo("%s %s: %s", prefix, error, msg)
}

// FirstLine prints a line with a leading arrow.
func (l Log) FirstLine(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printInfo("%s %s", prefix, msg)
}

// SubsequentLine prints indented output without the leading arrow.
func (l Log) SubsequentLine(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printInfo("%s%s", indent, msg)
}

// Warning prints a warning message
func (l Log) Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printInfo("%s %s: %s", prefix, warning, msg)
}

func (l Log) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printDebug("%s %s: %s", prefix, debug, msg)
}

func (l Log) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.printInfo("%s %s: %s", prefix, info, msg)
}

func (l Log) PrettyIdentity(v Identifiable) string {
	if v == nil {
		return ""
	}
	var sb strings.Builder
	name, description := v.Identity()
	if name != "" {
		sb.WriteString(name)
	}
	if description != "" {
		sb.WriteString(" - ")
		sb.WriteString(description)
	}
	return sb.String()
}

func (l Log) IsDebugEnabled() bool {
	l.Lock()
	defer l.Unlock()
	return l.debug != nil
}

func (l Log) printDebug(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()
	print(l.debug, format, args...)
}

func (l Log) printInfo(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()
	print(l.info, format, args...)
}

func print(w *bufio.Writer, format string, args ...interface{}) {
	if w != nil {
		msg := fmt.Sprintf(format, args...)
		_, _ = fmt.Fprintf(w, "%s\n", msg)
		_ = w.Flush()
	}
}
