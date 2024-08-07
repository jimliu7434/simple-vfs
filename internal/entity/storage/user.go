package storage

import (
	"fmt"
	"time"
)

// GetUser returns user data from the storage by username
// If the user does not exist, it returns an error
func (s *Storage) GetUser(username string) (*UserData, error) {
	if u, ok := s.users[username]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("user %s doesn't exist", username)
}

// CreateUser creates a new user in the storage
// If the user already exists, it returns an error
func (s *Storage) CreateUser(username string) error {
	if _, ok := s.users[username]; !ok {
		s.users[username] = &UserData{
			Name:      username,
			CreatedAt: time.Now(),
			folders:   map[string]*FolderData{},
		}
		return nil
	}
	return fmt.Errorf("user %s has already existed", username)
}

// UserData is the user data struct
// It contains the username, create time of the user
// and the folders of the user
type UserData struct {
	Name      string
	CreatedAt time.Time
	folders   map[string]*FolderData // if need multi-threading, use sync.Map
}
