// Package util is the package that contains the utility functions.
package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/gosuri/uitable"
)

// Info prints the message to the stdout
func Info(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf(format, args...)
}

// Error prints the message to the stderr
func Error(format string, args ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(os.Stderr, format, args...)
}

// Table prints the table to the stdout
func Table(rows [][]any) {
	table := uitable.New()
	table.Wrap = true
	table.MaxColWidth = 100

	for _, row := range rows {
		table.AddRow(row...)
	}
}
