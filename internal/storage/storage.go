package storage

import "fmt"

// Storage is the storage struct
type Storage struct {
	users map[string]*UserData
}

// ErrUserNotFound is returned when the user is not found in the storage
var ErrUserNotFound = fmt.Errorf("user not found")

// GetUser returns user data from the storage by username
func (s *Storage) GetUser(username string) (*UserData, error) {
	if u, ok := s.users[username]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}
