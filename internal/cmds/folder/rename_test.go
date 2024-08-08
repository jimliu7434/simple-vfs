package folder

import (
	"context"
	"fmt"
	"strings"
	"testing"

	Storage "simple-vfs/internal/entity/storage"

	"github.com/urfave/cli/v2"
)

func TestBeforeRename(t *testing.T) {
	testCases := []struct {
		name                string
		args                []string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "valid newfoldername",
			args:                []string{"user1", "folder1", "newfolder1"},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "invalid newfoldername",
			args:                []string{"user1", "folder1", "newfolder!1"},
			expectedError:       true,
			expectedErrContains: "foldername newfolder!1 contains invalid chars",
		},
	}

	app := &cli.App{
		Name: "test",
		Commands: []*cli.Command{
			{
				Name:   "rename",
				Before: BeforeRename,
				Action: func(_ *cli.Context) error { return nil },
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"test", "rename"}, tc.args...)

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

func TestActionRename(t *testing.T) {
	// prepare storage
	username := "user1"
	foldername := "folder1"
	foldername2 := "folder2"

	storage := Storage.New()
	storage.CreateUser(username)
	user, _ := storage.GetUser(username)
	user.CreateFolder(foldername, "")
	user.CreateFolder(fmt.Sprint("new", foldername2), "")

	testCases := []struct {
		name                string
		username            string
		foldername          string
		newfoldername       string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "user not exist",
			username:            "user0",
			foldername:          foldername,
			newfoldername:       fmt.Sprint("new", foldername),
			expectedError:       true,
			expectedErrContains: "user user0 doesn't exist",
		},
		{
			name:                "folder not exists",
			username:            username,
			foldername:          "folder0",
			newfoldername:       fmt.Sprint("new", "folder0"),
			expectedError:       true,
			expectedErrContains: "folder folder0 doesn't exist",
		},
		{
			name:                "newfoldername exists",
			username:            username,
			foldername:          foldername2,
			newfoldername:       fmt.Sprint("new", foldername2),
			expectedError:       true,
			expectedErrContains: "",
		},
		{
			name:                "rename ok",
			username:            username,
			foldername:          foldername,
			newfoldername:       fmt.Sprint("new", foldername),
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

			c.Context = context.WithValue(c.Context, argsKey, &renameArgs{
				username:      tc.username,
				foldername:    tc.foldername,
				newfoldername: tc.newfoldername,
			})

			err := ActionRename(c)
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
