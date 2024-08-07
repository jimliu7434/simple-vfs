package storage

// Storage is the storage struct
type Storage struct {
	users map[string]*UserData // if need multi-threading, use sync.Map
}
