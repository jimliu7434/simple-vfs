package storage

import "time"

// UserData is the user data struct
// It contains the username, create time of the user
// and the folders of the user
type UserData struct {
	Name      string
	CreatedAt time.Time
	folders   map[string]FolderData
}

// AddFolder adds a new folder to the user
// If the folder already exists, it returns an error
func (u *UserData) AddFolder(f FolderData) error {
	// TODO
	return nil
}

// DelFolder deletes a folder from the user
// It also deletes all the files in the folder
// If the folder does not exist, it returns an error
func (u *UserData) DelFolder(foldername string) error {
	// TODO
	return nil
}

// RenameFolder renames a folder of the user
// If the folder does not exist, it returns an error
func (u *UserData) RenameFolder(oldname, newname string) error {
	// TODO
	return nil
}

// ListFolders lists all folders of the user
func (u *UserData) ListFolders(foldername string, sortBy sortBy, sortType sortType) []FolderData {
	// TODO
	return []FolderData{}
}
