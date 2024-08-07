package file

import (
	"context"
	"fmt"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"

	"github.com/urfave/cli/v2"
)

type listArgs struct {
	username   string
	foldername string
	sortBy     Storage.SortBy
	sortType   Storage.SortType
}

func (args *listArgs) IsValid() error {
	// if !Storage.IsValidFoldername(args.foldername) {
	// 	return fmt.Errorf("Invalid foldername %s", args.foldername)
	// }

	if !Storage.IsValidSortType(string(args.sortType)) {
		return fmt.Errorf("Invalid sort type %s", args.sortType)
	}

	if args.sortBy != Storage.SortByName && args.sortBy != Storage.SortByTime {
		args.sortBy = _DefaultSortBy
	}

	if args.sortType != Storage.SortAsc && args.sortType != Storage.SortDesc {
		args.sortType = _DefaultSortType
	}
	return nil
}

var _DefaultSortBy = Storage.SortByName
var _DefaultSortType = Storage.SortAsc

// BeforeList is the command before the List command
func BeforeList(c *cli.Context) error {
	sortBy := Storage.SortByName
	sortType := Storage.SortAsc

	isSortByName := c.String("sort-name")
	isSortByTime := c.String("sort-created")

	if isSortByName != "" {
		sortBy = Storage.SortByName
		sortType = Storage.SortType(c.String("sort-name"))
	} else if isSortByTime != "" {
		sortBy = Storage.SortByTime
		sortType = Storage.SortType(c.String("sort-created"))
	}

	args := &listArgs{
		username:   c.Args().Get(0),
		foldername: c.Args().Get(1),
		sortBy:     sortBy,
		sortType:   sortType,
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionList is the action to list all files in a folder
func ActionList(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*listArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(args.foldername)
	if err != nil {
		return err
	}

	files := folder.ListFiles(args.sortBy, args.sortType)

	// folder doesn't have any file warning
	if len(files) == 0 {
		logger.Warn("folder %s is empty", args.foldername)
		return nil
	}

	// print result
	printRows := make([][]any, 0, len(files))
	printRows = append(printRows, []any{"File", "Desc", "Created At", "Folder", "Owner"})
	for _, f := range files {
		printRows = append(printRows,
			[]any{
				f.Name,
				f.Description,
				f.CreatedAt.Format("2006-01-02 15:04:05"),
				folder.Name,
				user.Name,
			},
		)
	}
	logger.Table(printRows)

	return nil
}
