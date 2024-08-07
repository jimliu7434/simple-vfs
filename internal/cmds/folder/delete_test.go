package folder

import (
	"context"
	"strings"
	"testing"

	Storage "simple-vfs/internal/entity/storage"

	"github.com/urfave/cli/v2"
)

func TestBeforeDelete(t *testing.T) {
	testCases := []struct {
		name                string
		args                []string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "valid foldername",
			args:                []string{"user1", "folder1"},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "invalid foldername",
			args:                []string{"user1", "folder!1"},
			expectedError:       true,
			expectedErrContains: "Invalid foldername folder!1",
		},
	}

	app := &cli.App{
		Name: "test",
		Commands: []*cli.Command{
			{
				Name:   "delete",
				Before: BeforeDelete,
				Action: func(_ *cli.Context) error { return nil },
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"test", "delete"}, tc.args...)

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

func TestActionDelete(t *testing.T) {
	// prepare storage
	username := "user1"
	foldername := "folder1"

	storage := Storage.New()
	storage.CreateUser(username)
	user, _ := storage.GetUser(username)
	user.CreateFolder(foldername, "")

	testCases := []struct {
		name                string
		username            string
		foldername          string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "user not exist",
			username:            "user0",
			foldername:          foldername,
			expectedError:       true,
			expectedErrContains: "user user0 doesn't exist",
		},
		{
			name:                "folder not exists",
			username:            username,
			foldername:          "folder0",
			expectedError:       true,
			expectedErrContains: "folder folder0 doesn't exist",
		},
		{
			name:                "folder exists",
			username:            username,
			foldername:          foldername,
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

			c.Context = context.WithValue(c.Context, argsKey, &deleteArgs{
				username:   tc.username,
				foldername: tc.foldername,
			})

			err := ActionDelete(c)
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
