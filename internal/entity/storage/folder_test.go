package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFolder(t *testing.T) {
	user := UserData{
		Name:      "test",
		CreatedAt: time.Now(),
		folders: map[string]*FolderData{
			"folder1": {
				Name:        "folder1",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
		},
	}

	t.Run("folder not exist", func(t *testing.T) {
		folder, err := user.GetFolder("notexist")
		if err == nil {
			t.Errorf("GetFolder() failed, err is nil")
		}
		if folder != nil {
			t.Errorf("GetFolder() failed, folder is not nil")
		}
		assert.Contains(t, "folder notexist doesn't exist", err.Error())
	})

	t.Run("folder exist", func(t *testing.T) {
		folder, err := user.GetFolder("folder1")
		if err != nil {
			t.Errorf("GetFolder() failed, err is not nil %v", err)
		}
		if folder == nil {
			t.Errorf("GetFolder() failed, folder is nil")
		}
		assert.Equal(t, "folder1", folder.Name)
	})
}

func TestCreateFolder(t *testing.T) {
	user := UserData{
		Name:      "test",
		CreatedAt: time.Now(),
		folders: map[string]*FolderData{
			"folder1": {
				Name:        "folder1",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
		},
	}

	t.Run("folder not exist", func(t *testing.T) {
		newFolder := "folder2"
		err := user.CreateFolder(newFolder, "desc")
		if err != nil {
			t.Errorf("CreateFolder() failed, err is not nil %v", err)
		}
	})

	t.Run("folder exist", func(t *testing.T) {
		err := user.CreateFolder("folder1", "desc")
		if err == nil {
			t.Errorf("CreateFolder() failed, err is nil")
		}
		assert.Contains(t, "folder folder1 has already exist", err.Error())
	})
}

func TestDelFolder(t *testing.T) {
	user := UserData{
		Name:      "test",
		CreatedAt: time.Now(),
		folders: map[string]*FolderData{
			"folder1": {
				Name:        "folder1",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
		},
	}

	t.Run("folder not exist", func(t *testing.T) {
		err := user.DelFolder("notexist")
		if err == nil {
			t.Errorf("DelFolder() failed, err is nil")
		}
		assert.Contains(t, "folder notexist doesn't exist", err.Error())
	})

	t.Run("folder exist", func(t *testing.T) {
		err := user.DelFolder("folder1")
		if err != nil {
			t.Errorf("DelFolder() failed, err is not nil %v", err)
		}
	})
}

func TestRenameFolder(t *testing.T) {
	user := UserData{
		Name:      "test",
		CreatedAt: time.Now(),
		folders: map[string]*FolderData{
			"folder1": {
				Name:        "folder1",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
		},
	}

	t.Run("old folder not exist", func(t *testing.T) {
		err := user.RenameFolder("notexist", "newfolder")
		if err == nil {
			t.Errorf("RenameFolder() failed, err is nil")
		}
		assert.Contains(t, "folder notexist doesn't exist", err.Error())
	})

	t.Run("new folder exist", func(t *testing.T) {
		err := user.RenameFolder("folder1", "folder1")
		if err == nil {
			t.Errorf("RenameFolder() failed, err is nil")
		}
		assert.Contains(t, "folder folder1 has already exist", err.Error())
	})

	t.Run("old folder exist", func(t *testing.T) {
		err := user.RenameFolder("folder1", "newfolder")
		if err != nil {
			t.Errorf("RenameFolder() failed, err is not nil %v", err)
		}
	})
}

func TestListFolders(t *testing.T) {
	user := UserData{
		Name:      "test",
		CreatedAt: time.Now(),
		folders: map[string]*FolderData{
			"folder1": {
				Name:        "folder1",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
			"folder2": {
				Name:        "folder2",
				Description: "desc",
				CreatedAt:   time.Now(),
				files:       map[string]*FileData{},
			},
		},
	}

	t.Run("sort by name asc", func(t *testing.T) {
		folders := user.ListFolders(SortByName, SortAsc)
		assert.Len(t, folders, 2)
		assert.Equal(t, "folder1", folders[0].Name)
		assert.Equal(t, "folder2", folders[1].Name)
	})

	t.Run("sort by name desc", func(t *testing.T) {
		folders := user.ListFolders(SortByName, SortDesc)
		assert.Len(t, folders, 2)
		assert.Equal(t, "folder2", folders[0].Name)
		assert.Equal(t, "folder1", folders[1].Name)
	})

	t.Run("sort by time asc", func(t *testing.T) {
		folders := user.ListFolders(SortByTime, SortAsc)
		assert.Len(t, folders, 2)
	})

	t.Run("sort by time desc", func(t *testing.T) {
		folders := user.ListFolders(SortByTime, SortDesc)
		assert.Len(t, folders, 2)
	})
}
