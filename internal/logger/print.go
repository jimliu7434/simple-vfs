// Package logger is the package that contains the utility functions.
package logger

import (
	"fmt"
	"strings"

	"github.com/gosuri/uitable"
)

// Info prints the message to the stdout
func Info(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(outWriter, format, args...)
}

// Warn prints the message to the stderr, with a prefix "Warning: "
func Warn(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if !strings.HasPrefix(format, "Warning: ") {
		format = "Warning: " + format
	}
	fmt.Fprintf(outWriter, format, args...)
}

// Error prints the message to the stderr, with a prefix "Error: "
func Error(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if !strings.HasPrefix(format, "Error: ") {
		format = "Error: " + format
	}
	fmt.Fprintf(errWriter, format, args...)
}

// Table prints the table to the stdout
func Table(rows [][]any) {
	table := uitable.New()
	table.Wrap = true
	table.MaxColWidth = 100

	for _, row := range rows {
		table.AddRow(row...)
	}

	fmt.Fprintln(outWriter, table)
}
