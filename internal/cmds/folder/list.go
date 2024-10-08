package folder

import (
	"context"
	"fmt"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"

	"github.com/urfave/cli/v2"
)

type listArgs struct {
	username string
	sortBy   Storage.SortBy
	sortType Storage.SortType
}

func (args *listArgs) IsValid() error {
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

	if c.String("sort-name") != "" {
		sortBy = Storage.SortByName
		sortType = Storage.SortType(c.String("sort-name"))
	} else if c.String("sort-created") != "" {
		sortBy = Storage.SortByTime
		sortType = Storage.SortType(c.String("sort-created"))
	}

	args := &listArgs{
		username: c.Args().Get(0),
		sortBy:   sortBy,
		sortType: sortType,
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionList is the action to list all folders
func ActionList(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*listArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	folders := user.ListFolders(args.sortBy, args.sortType)

	if len(folders) == 0 {
		logger.Warn("user %s doesn't have any folder", args.username)
		return nil
	}

	// print result
	printRows := make([][]any, 0, len(folders))
	printRows = append(printRows, []any{"Folder", "Desc", "Created At", "Owner"})
	for _, f := range folders {
		printRows = append(printRows,
			[]any{
				f.Name,
				f.Description,
				f.CreatedAt.Format("2006-01-02 15:04:05"),
				user.Name,
			},
		)
	}
	logger.Table(printRows)

	return nil
}
