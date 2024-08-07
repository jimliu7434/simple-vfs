package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.Equal(t, outWriter, os.Stdout)
	assert.Equal(t, errWriter, os.Stderr)
}

func TestSetOutWriter(t *testing.T) {
	output := &bytes.Buffer{}
	SetOutWriter(output)
	assert.Equal(t, outWriter, output)
}

func TestSetErrWriter(t *testing.T) {
	output := &bytes.Buffer{}
	SetErrWriter(output)
	assert.Equal(t, errWriter, output)
}
