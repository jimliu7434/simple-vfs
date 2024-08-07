package storage

import "time"

// FolderData is the folder data struct
// It contains the folder name, create time of the folder
// and the files of the folder
type FolderData struct {
	Name      string
	CreatedAt time.Time
	files     map[string]FileData
}

// AddFile adds a new File to the folder
// If the File already exists, it returns an error
func (u *FolderData) AddFile(f FileData) error {
	// TODO
	return nil
}

// DelFile deletes a File from the folder
// If the File does not exist, it returns an error
func (u *FolderData) DelFile(filename string) error {
	// TODO
	return nil
}

// ListFiles lists all files of the folder
func (u *FolderData) ListFiles(sortBy sortBy, sortType sortType) []FolderData {
	// TODO
	return []FolderData{}
}
