package storage

import "time"

// FileData is the file data struct
// It contains the file name and create time of the file
type FileData struct {
	Name      string
	CreatedAt time.Time
}
