package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	storage := Storage{
		users: map[string]*UserData{
			"test": {
				Name:      "test",
				CreatedAt: time.Now(),
				folders:   map[string]*FolderData{},
			},
		},
	}

	t.Run("user not exist", func(t *testing.T) {
		user, err := storage.GetUser("notexist")
		if err == nil {
			t.Errorf("GetUser() failed, err is nil")
		}
		if user != nil {
			t.Errorf("GetUser() failed, user is not nil")
		}
		assert.Contains(t, "user notexist doesn't exist", err.Error())
	})

	t.Run("user exist", func(t *testing.T) {
		user, err := storage.GetUser("test")
		if err != nil {
			t.Errorf("GetUser() failed, err is not nil %v", err)
		}
		if user == nil {
			t.Errorf("GetUser() failed, user is nil")
		}
		assert.Equal(t, "test", user.Name)
	})
}

func TestCreateUser(t *testing.T) {
	storage := Storage{
		users: map[string]*UserData{
			"test": {
				Name:      "test",
				CreatedAt: time.Now(),
				folders:   map[string]*FolderData{},
			},
		},
	}

	t.Run("user not exist", func(t *testing.T) {
		newUser := "another"
		err := storage.CreateUser(newUser)
		if err != nil {
			t.Errorf("CreateUser() failed, err is not nil %v", err)
		}
		assert.Len(t, storage.users, 2)
		if user, ok := storage.users[newUser]; !ok {
			t.Errorf("CreateUser() failed, user %s not found", newUser)
		} else {
			assert.Equal(t, newUser, user.Name)
		}
	})

	t.Run("user exist", func(t *testing.T) {
		err := storage.CreateUser("test")
		if err == nil {
			t.Errorf("CreateUser() failed, err is nil")
		}
		assert.Contains(t, "user test has already existed", err.Error())
	})
}
