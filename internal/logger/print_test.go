package logger_test

import (
	"bytes"
	"simple-vfs/internal/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	output := bytes.Buffer{}
	logger.SetOutWriter(&output)

	logger.Info("This is a test")
	assert.Equal(t, "This is a test\n", output.String())
}

func TestWarn(t *testing.T) {
	output := bytes.Buffer{}
	logger.SetOutWriter(&output)

	logger.Warn("This is a test")
	assert.Equal(t, "Warning: This is a test\n", output.String())
}

func TestError(t *testing.T) {
	output := bytes.Buffer{}
	logger.SetErrWriter(&output)

	logger.Error("This is a test")
	assert.Equal(t, "Error: This is a test\n", output.String())
}

func TestTable(t *testing.T) {
	output := bytes.Buffer{}
	logger.SetOutWriter(&output)

	logger.Table([][]any{
		{"a", "b", "c"},
		{"d", "e", "f"},
	})
	assert.Equal(t, "a\tb\tc\nd\te\tf\n", output.String())
}
