package storage

import (
	"fmt"
	"time"
)

// GetFolder returns a folder of the user by foldername
// If the folder does not exist, it returns an error
func (u *UserData) GetFolder(foldername string) (*FolderData, error) {
	if f, ok := u.folders[foldername]; ok {
		return f, nil
	}
	return nil, fmt.Errorf("folder %s doesn't exist", foldername)
}

// CreateFolder adds a new folder to the user
// If the folder already exists, it returns an error
func (u *UserData) CreateFolder(foldername, desc string) error {
	if _, ok := u.folders[foldername]; !ok {
		u.folders[foldername] = &FolderData{
			Name:        foldername,
			Description: desc,
			CreatedAt:   time.Now(),
			files:       map[string]*FileData{},
			Owner:       u,
		}
		return nil
	}

	return fmt.Errorf("folder %s has already existed", foldername)
}

// DelFolder deletes a folder from the user
// It also deletes all the files in the folder
// If the folder does not exist, it returns an error
func (u *UserData) DelFolder(foldername string) error {
	if _, ok := u.folders[foldername]; ok {
		delete(u.folders, foldername)
		return nil
	}
	return fmt.Errorf("folder %s doesn't exist", foldername)
}

// RenameFolder renames a folder of the user
// If the old folder name does not exist, it returns an error
// If the new folder name already exists, it returns an error
func (u *UserData) RenameFolder(oldname, newname string) error {
	if _, ok := u.folders[newname]; ok {
		return fmt.Errorf("folder %s has already existed", newname)
	}

	if _, ok := u.folders[oldname]; ok {
		u.folders[newname] = u.folders[oldname]
		u.folders[newname].Name = newname
		delete(u.folders, oldname)
		return nil
	}

	return fmt.Errorf("folder %s doesn't exist", oldname)
}

// ListFolders lists all folders of the user
// It returns the folders sorted by the sortBy and sortType
func (u *UserData) ListFolders(sortBy SortBy, sortType SortType) []*FolderData {
	rtn := make([]*FolderData, 0, len(u.folders))
	for _, v := range u.folders {
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

// FolderData is the folder data struct
// It contains the folder name, description, create time of the folder
// and the files of the folder
type FolderData struct {
	Name        string
	Description string
	CreatedAt   time.Time
	files       map[string]*FileData // if need multi-threading, use sync.Map
	Owner       *UserData
}

// GetName returns the name of the folder
func (f *FolderData) GetName() string {
	return f.Name
}

// GetCreatedAt returns the create time of the folder
func (f *FolderData) GetCreatedAt() time.Time {
	return f.CreatedAt
}
