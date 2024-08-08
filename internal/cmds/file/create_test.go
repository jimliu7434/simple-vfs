package file

import (
	"context"
	"fmt"
	"strings"
	"testing"

	Storage "simple-vfs/internal/entity/storage"

	"github.com/urfave/cli/v2"
)

func TestBeforeCreate(t *testing.T) {
	testCases := []struct {
		name                string
		args                []string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "valid filename",
			args:                []string{"user1", "folder1", "file1"},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "invalid filename",
			args:                []string{"user1", "folder1", "file!1"},
			expectedError:       true,
			expectedErrContains: "filename file!1 contains invalid chars",
		},
	}

	app := &cli.App{
		Name: "test",
		Commands: []*cli.Command{
			{
				Name:   "create",
				Before: BeforeCreate,
				Action: func(_ *cli.Context) error { return nil },
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"test", "create"}, tc.args...)

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

func TestActionCreate(t *testing.T) {
	// prepare storage
	username := "user1"
	foldername := "folder1"
	filename := "file1"

	storage := Storage.New()
	storage.CreateUser(username)
	user, _ := storage.GetUser(username)
	user.CreateFolder(foldername, "")
	folder, _ := user.GetFolder(foldername)
	folder.CreateFile(filename, "")

	testCases := []struct {
		name                string
		username            string
		foldername          string
		filename            string
		description         string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "user not exist",
			username:            "user0",
			foldername:          foldername,
			filename:            filename,
			description:         "",
			expectedError:       true,
			expectedErrContains: "user user0 doesn't exist",
		},
		{
			name:                "folder not exist",
			username:            username,
			foldername:          "folder0",
			filename:            filename,
			description:         "",
			expectedError:       true,
			expectedErrContains: "folder folder0 doesn't exist",
		},
		{
			name:                "file exist",
			username:            username,
			foldername:          foldername,
			filename:            filename,
			description:         "",
			expectedError:       true,
			expectedErrContains: fmt.Sprintf("file %s has already existed", filename),
		},
		{
			name:                "create file, description without space",
			username:            username,
			foldername:          foldername,
			filename:            "file2",
			description:         "my-description",
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "create file, description with space",
			username:            username,
			foldername:          foldername,
			filename:            "file3",
			description:         "my description",
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
				username:    tc.username,
				foldername:  tc.foldername,
				filename:    tc.filename,
				description: tc.description,
			})

			err := ActionCreate(c)
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

				// check file exist
				user, _ := storage.GetUser(tc.username)
				folder, _ := user.GetFolder(tc.foldername)
				files := folder.ListFiles(Storage.SortByName, Storage.SortAsc)
				var file *Storage.FileData
				for _, f := range files {
					if f.Name == tc.filename {
						file = f
						break
					}
				}
				if file == nil {
					t.Errorf("expected file %s exist, got nil", tc.filename)
				}

				if file.Description != tc.description {
					t.Errorf("expected file description %s, got %s", tc.description, file.Description)
				}
			}
		})
	}
}
