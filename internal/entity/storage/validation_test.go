package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUsername(t *testing.T) {
	t.Run("valid username", func(t *testing.T) {
		assert.True(t, IsValidUsername("test"))
	})

	t.Run("len < 3", func(t *testing.T) {
		assert.False(t, IsValidUsername("t"))
	})

	t.Run("len > 20", func(t *testing.T) {
		assert.False(t, IsValidUsername("abcdef1234567890000000000000"))
	})

	t.Run("invalid char", func(t *testing.T) {
		assert.False(t, IsValidUsername("test!"))
	})
}

func TestIsValidFoldername(t *testing.T) {
	t.Run("valid folder name", func(t *testing.T) {
		assert.True(t, IsValidFoldername("test"))
	})

	t.Run("len < 1", func(t *testing.T) {
		assert.False(t, IsValidFoldername(""))
	})

	t.Run("len > 50", func(t *testing.T) {
		assert.False(t, IsValidFoldername("abcdef1234567890000000000000alsdjflakjsdlfkjasldkjfalksdjflaksdjfasdfasdf"))
	})

	t.Run("invalid char", func(t *testing.T) {
		assert.False(t, IsValidFoldername("test;!@$$%"))
	})
}

func TestIsValidFilename(t *testing.T) {
	t.Run("valid file name", func(t *testing.T) {
		assert.True(t, IsValidFilename("test"))
	})

	t.Run("len < 1", func(t *testing.T) {
		assert.False(t, IsValidFilename(""))
	})

	t.Run("len > 50", func(t *testing.T) {
		assert.False(t, IsValidFilename("abcdef1234567890000000000000alsdjflakjsdlfkjasldkjfalksdjflaksdjfasdfasdf"))
	})

	t.Run("invalid char", func(t *testing.T) {
		assert.False(t, IsValidFilename("test;!@$$%"))
	})
}

func TestIsValidSortType(t *testing.T) {
	t.Run("valid sort type", func(t *testing.T) {
		assert.True(t, IsValidSortType("asc"))
		assert.True(t, IsValidSortType("desc"))
		assert.True(t, IsValidSortType("ASC"))
		assert.True(t, IsValidSortType("DESC"))
		assert.True(t, IsValidSortType("aSc"))
		assert.True(t, IsValidSortType("dESc"))
	})

	t.Run("invalid sort type", func(t *testing.T) {
		assert.False(t, IsValidSortType("invalid"))
	})
}
