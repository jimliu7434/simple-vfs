package user

import (
	"context"
	"fmt"
	"strings"
	"testing"

	Storage "simple-vfs/internal/entity/storage"

	"github.com/urfave/cli/v2"
)

func TestBeforeRegister(t *testing.T) {
	testCases := []struct {
		name                string
		args                []string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "valid username",
			args:                []string{"user1"},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "invalid username",
			args:                []string{"user!1"},
			expectedError:       true,
			expectedErrContains: "username user!1 contains invalid chars",
		},
	}

	app := &cli.App{
		Name: "test",
		Commands: []*cli.Command{
			{
				Name:   "register",
				Before: BeforeRegister,
				Action: func(_ *cli.Context) error { return nil },
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"test", "register"}, tc.args...)

			err := app.RunContext(context.Background(), args)

			if tc.expectedError {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if !strings.Contains(err.Error(), tc.expectedErrContains) {
					t.Errorf("expected error contains %s, got %s", tc.expectedErrContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
			}
		})
	}
}

func TestActionRegister(t *testing.T) {
	// prepare storage
	username := "user0"

	storage := Storage.New()
	storage.CreateUser(username)

	testCases := []struct {
		name                string
		username            string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "user existed",
			username:            username,
			expectedError:       true,
			expectedErrContains: fmt.Sprintf("user %s has already existed", username),
		},
		{
			name:                "create user",
			username:            "user100",
			expectedError:       false,
			expectedErrContains: "",
		},
	}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{
		"storage": &storage,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := cli.NewContext(app, nil, nil)

			c.Context = context.WithValue(c.Context, argsKey, &createArgs{
				username: tc.username,
			})

			err := ActionRegister(c)
			if tc.expectedError {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if !strings.Contains(err.Error(), tc.expectedErrContains) {
					t.Errorf("expected error contains %s, got %s", tc.expectedErrContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}

				// check user exist
				_, err := storage.GetUser(tc.username)

				if err != nil {
					t.Errorf("expected user exist, got %s", err.Error())
				}
			}
		})
	}
}
