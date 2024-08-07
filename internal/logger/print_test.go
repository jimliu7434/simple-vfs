package logger_test

import (
	"bytes"
	"log"
	"os"
	"simple-vfs/internal/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	output := captureStdOut(func() {
		logger.Info("This is a test")
	})
	assert.Equal(t, "This is a test\n", output)
}

func captureStdOut(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
