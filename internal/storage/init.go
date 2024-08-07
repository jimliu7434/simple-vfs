package storage

type sortBy string
type sortType string

// SortByName is the sort by name (foldername or filename)
const SortByName sortBy = "name"

// SortByTime is the sort by create time
const SortByTime sortBy = "time"

// SortAsc is the sort in ascending order
const SortAsc sortType = "asc"

// SortDesc is the sort in descending order
const SortDesc sortType = "desc"

// New creates a new empty storage
func New() Storage {
	return Storage{
		users: map[string]*UserData{},
	}
}
