package storage

import (
	"fmt"
	"time"
)

// CreateFile adds a new File to the folder
// If the File already exists, it returns an error
func (u *FolderData) CreateFile(filename, desc string) error {
	if _, ok := u.files[filename]; !ok {
		u.files[filename] = &FileData{
			Name:        filename,
			Description: desc,
			CreatedAt:   time.Now(),
			Folder:      u,
		}
		return nil
	}
	return fmt.Errorf("file %s has already existed", filename)
}

// DelFile deletes a File from the folder
// If the File does not exist, it returns an error
func (u *FolderData) DelFile(filename string) error {
	if _, ok := u.files[filename]; ok {
		delete(u.files, filename)
		return nil
	}
	return fmt.Errorf("file %s doesn't exist", filename)
}

// ListFiles lists all files of the folder
func (u *FolderData) ListFiles(sortBy SortBy, sortType SortType) []*FileData {
	rtn := make([]*FileData, 0, len(u.files))
	for _, v := range u.files {
		rtn = append(rtn, v)
	}

	doReverse := false
	if sortType == SortDesc {
		doReverse = true
	}

	if sortBy == SortByName {
		sortAscByName(rtn, doReverse)
	} else {
		sortAscByTime(rtn, doReverse)
	}
	return rtn
}

// FileData is the file data struct
// It contains the file name, description and create time of the file
type FileData struct {
	Name        string
	Description string
	CreatedAt   time.Time
	Folder      *FolderData
}

// GetName returns the name of the file
func (f *FileData) GetName() string {
	return f.Name
}

// GetCreatedAt returns the create time of the file
func (f *FileData) GetCreatedAt() time.Time {
	return f.CreatedAt
}
