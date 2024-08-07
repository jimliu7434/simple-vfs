package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	folder := FolderData{
		Name:        "folder1",
		Description: "desc",
		CreatedAt:   time.Now(),
		files: map[string]*FileData{
			"file1": {
				Name:        "file1",
				Description: "desc",
				CreatedAt:   time.Now(),
			},
		},
	}

	t.Run("file not exist", func(t *testing.T) {
		err := folder.CreateFile("file2", "desc")
		if err != nil {
			t.Errorf("CreateFile() failed, err is not nil %v", err)
		}
	})

	t.Run("file exist", func(t *testing.T) {
		err := folder.CreateFile("file1", "desc")
		if err == nil {
			t.Errorf("CreateFile() failed, err is nil")
		}
		assert.Contains(t, "file file1 has already existed", err.Error())
	})
}

func TestDelFile(t *testing.T) {
	folder := FolderData{
		Name:        "folder1",
		Description: "desc",
		CreatedAt:   time.Now(),
		files: map[string]*FileData{
			"file1": {
				Name:        "file1",
				Description: "desc",
				CreatedAt:   time.Now(),
			},
		},
	}

	t.Run("file exist", func(t *testing.T) {
		err := folder.DelFile("file1")
		if err != nil {
			t.Errorf("DelFile() failed, err is not nil %v", err)
		}
	})

	t.Run("file not exist", func(t *testing.T) {
		err := folder.DelFile("file2")
		if err == nil {
			t.Errorf("DelFile() failed, err is nil")
		}
		assert.Contains(t, "file file2 doesn't exist", err.Error())
	})
}

func TestListFiles(t *testing.T) {
	folder := FolderData{
		Name:        "folder1",
		Description: "desc",
		CreatedAt:   time.Now(),
		files: map[string]*FileData{
			"file1": {
				Name:        "file1",
				Description: "desc",
				CreatedAt:   time.Now(),
			},
			"file2": {
				Name:        "file2",
				Description: "desc",
				CreatedAt:   time.Now(),
			},
		},
	}

	t.Run("sort by name asc", func(t *testing.T) {
		files := folder.ListFiles(SortByName, SortAsc)
		assert.Len(t, files, 2)
		assert.Equal(t, "file1", files[0].Name)
		assert.Equal(t, "file2", files[1].Name)
	})

	t.Run("sort by name desc", func(t *testing.T) {
		files := folder.ListFiles(SortByName, SortDesc)
		assert.Len(t, files, 2)
		assert.Equal(t, "file2", files[0].Name)
		assert.Equal(t, "file1", files[1].Name)
	})

	t.Run("sort by time asc", func(t *testing.T) {
		files := folder.ListFiles(SortByTime, SortAsc)
		assert.Len(t, files, 2)
		assert.Equal(t, "file1", files[0].Name)
		assert.Equal(t, "file2", files[1].Name)
	})

	t.Run("sort by time desc", func(t *testing.T) {
		files := folder.ListFiles(SortByTime, SortDesc)
		assert.Len(t, files, 2)
		assert.Equal(t, "file2", files[0].Name)
		assert.Equal(t, "file1", files[1].Name)
	})
}
