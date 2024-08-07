// Package storage is the package that contains the storage entity and its methods.
package storage

import "time"

// SortBy is the sort by type in one of the following: name, time
type SortBy string

// SortType is the sort type in one of the following: asc, desc
type SortType string

// SortByName is the sort by name (foldername or filename)
const SortByName SortBy = "name"

// SortByTime is the sort by create time
const SortByTime SortBy = "time"

// SortAsc is the sort in ascending order
const SortAsc SortType = "asc"

// SortDesc is the sort in descending order
const SortDesc SortType = "desc"

// New creates a new empty storage
func New() Storage {
	return Storage{
		users: map[string]*UserData{},
	}
}

// ISortable is the interface that any entity that can be sorted should implement
type ISortable interface {
	GetName() string
	GetCreatedAt() time.Time
}

var _ ISortable = &FolderData{}
var _ ISortable = &FileData{}
