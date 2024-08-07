package file

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestBeforeList(t *testing.T) {
	testCases := []struct {
		name                string
		args                []string
		flags               [][]string
		expectedError       bool
		expectedErrContains string
	}{
		{
			name:                "sort by time desc",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-created", "desc"}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "sort by time asc",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-created", "asc"}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "sort by time, empty sort type",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-created", ""}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "sort by name desc",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-name", "desc"}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "sort by name asc",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-name", "asc"}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "sort by name, empty sort type",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-name", ""}},
			expectedError:       false,
			expectedErrContains: "",
		},
		{
			name:                "invalid sort type",
			args:                []string{"user1", "folder1"},
			flags:               [][]string{{"sort-name", "mysort"}},
			expectedError:       true,
			expectedErrContains: "Invalid sort type mysort",
		},
	}

	app := &cli.App{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"test", "list"}, tc.args...)

			// add flag into cli.Context
			set := flag.NewFlagSet("test", 0)
			set.String("sort-name", "", "")
			set.String("sort-created", "", "")
			set.Parse(args)
			c := cli.NewContext(app, set, nil)
			for _, flag := range tc.flags {
				c.Set(flag[0], flag[1])
			}

			err := BeforeList(c)

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

func TestActionList(t *testing.T) {
	// prepare storage
	username := "user1"
	foldername1 := "folder1"
	filename1 := "file1"
	filename2 := "file2"

	storage := Storage.New()
	storage.CreateUser(username)
	user, _ := storage.GetUser(username)
	user.CreateFolder(foldername1, "")
	folder, _ := user.GetFolder(foldername1)
	folder.CreateFile(filename1, "")
	time.Sleep(time.Second)
	folder.CreateFile(filename2, "")

	testCases := []struct {
		name                string
		username            string
		foldername          string
		sortBy              Storage.SortBy
		sortType            Storage.SortType
		expectedError       bool
		expectedErrContains string
		resultSort          []string
	}{
		{
			name:                "user not exist",
			username:            "user0",
			foldername:          foldername1,
			sortBy:              Storage.SortByName,
			sortType:            Storage.SortAsc,
			expectedError:       true,
			expectedErrContains: "user user0 doesn't exist",
			resultSort:          []string{},
		},
		{
			name:                "folder not exist",
			username:            username,
			foldername:          "folder0",
			sortBy:              Storage.SortByName,
			sortType:            Storage.SortAsc,
			expectedError:       true,
			expectedErrContains: "folder folder0 doesn't exist",
			resultSort:          []string{},
		},
		{
			name:                "sort by name asc",
			username:            username,
			foldername:          foldername1,
			sortBy:              Storage.SortByName,
			sortType:            Storage.SortAsc,
			expectedError:       false,
			expectedErrContains: "",
			resultSort:          []string{filename1, filename2},
		},
		{
			name:                "sort by name desc",
			username:            username,
			foldername:          foldername1,
			sortBy:              Storage.SortByName,
			sortType:            Storage.SortDesc,
			expectedError:       false,
			expectedErrContains: "",
			resultSort:          []string{filename2, filename1},
		},
		{
			name:                "sort by time asc",
			username:            username,
			foldername:          foldername1,
			sortBy:              Storage.SortByTime,
			sortType:            Storage.SortAsc,
			expectedError:       false,
			expectedErrContains: "",
			resultSort:          []string{filename1, filename2},
		},
		{
			name:                "sort by time desc",
			username:            username,
			foldername:          foldername1,
			sortBy:              Storage.SortByTime,
			sortType:            Storage.SortDesc,
			expectedError:       false,
			expectedErrContains: "",
			resultSort:          []string{filename2, filename1},
		},
	}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{
		"storage": &storage,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := cli.NewContext(app, nil, nil)

			c.Context = context.WithValue(c.Context, argsKey, &listArgs{
				username:   tc.username,
				foldername: tc.foldername,
				sortBy:     tc.sortBy,
				sortType:   tc.sortType,
			})

			var outWriter = &bytes.Buffer{}
			logger.SetOutWriter(outWriter)

			err := ActionList(c)
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

				// check stdout
				out := outWriter.String()
				outLines := strings.Split(out, "\n")
				// remove header
				outLines = outLines[1:]
				// remove empty line
				outLines = outLines[:len(outLines)-1]

				for i, line := range outLines {
					if i >= len(tc.resultSort) {
						break
					}
					assert.Contains(t, line, tc.resultSort[i])
				}
			}
		})
	}
}

func TestActionListWithNoFile(t *testing.T) {
	// prepare storage
	username := "user1"
	foldername1 := "folder1"

	storage := Storage.New()
	storage.CreateUser(username)
	user, _ := storage.GetUser(username)
	user.CreateFolder(foldername1, "")

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{
		"storage": &storage,
	}

	c := cli.NewContext(app, nil, nil)

	c.Context = context.WithValue(c.Context, argsKey, &listArgs{
		username:   username,
		foldername: foldername1,
		sortBy:     Storage.SortByName,
		sortType:   Storage.SortAsc,
	})

	var outWriter = &bytes.Buffer{}
	logger.SetOutWriter(outWriter)

	err := ActionList(c)
	assert.Nil(t, err)
	// check stdout
	assert.Contains(t, outWriter.String(), fmt.Sprintf("folder %s is empty", foldername1))
}
