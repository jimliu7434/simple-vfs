package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_sortAscByName(t *testing.T) {
	folders := []*FolderData{
		{
			Name:        "folder1",
			Description: "desc",
			CreatedAt:   time.Now(),
			files:       map[string]*FileData{},
		},
		{
			Name:        "folder2",
			Description: "desc",
			CreatedAt:   time.Now(),
			files:       map[string]*FileData{},
		},
	}

	t.Run("sort asc by name", func(t *testing.T) {
		sortAscByName(folders, false)

		assert.Equal(t, "folder1", folders[0].Name)
		assert.Equal(t, "folder2", folders[1].Name)
	})

	t.Run("sort desc by name", func(t *testing.T) {
		sortAscByName(folders, true)

		assert.Equal(t, "folder2", folders[0].Name)
		assert.Equal(t, "folder1", folders[1].Name)
	})
}

func Test_sortAscByTime(t *testing.T) {
	folders := []*FolderData{
		{
			Name:        "folder3",
			Description: "desc",
			CreatedAt:   time.Now().Add(-time.Hour),
			files:       map[string]*FileData{},
		},
		{
			Name:        "folder4",
			Description: "desc",
			CreatedAt:   time.Now(),
			files:       map[string]*FileData{},
		},
	}

	t.Run("sort asc by time", func(t *testing.T) {
		sortAscByTime(folders, false)

		assert.Equal(t, "folder3", folders[0].Name)
		assert.Equal(t, "folder4", folders[1].Name)
	})

	t.Run("sort desc by time", func(t *testing.T) {
		sortAscByTime(folders, true)

		assert.Equal(t, "folder4", folders[0].Name)
		assert.Equal(t, "folder3", folders[1].Name)
	})
}
