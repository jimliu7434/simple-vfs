package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	storage := New()
	assert.NotNil(t, storage)
	assert.NotNil(t, storage.users)
}
