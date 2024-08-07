package file

import (
	"context"
	"fmt"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"

	"github.com/urfave/cli/v2"
)

type deleteArgs struct {
	username   string
	foldername string
	filename   string
}

func (args *deleteArgs) IsValid() error {
	// if !Storage.IsValidFoldername(args.foldername) {
	// 	return fmt.Errorf("Invalid foldername %s", args.foldername)
	// }

	if !Storage.IsValidFilename(args.filename) {
		return fmt.Errorf("Invalid filename %s", args.foldername)
	}

	return nil
}

// BeforeDelete is the command before the Delete command
func BeforeDelete(c *cli.Context) error {
	args := &deleteArgs{
		username:   c.Args().Get(0),
		foldername: c.Args().Get(1),
		filename:   c.Args().Get(2),
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)
	return nil
}

// ActionDelete is the action to delete a file
func ActionDelete(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*deleteArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(args.foldername)
	if err != nil {
		return err
	}

	if err := folder.DelFile(args.filename); err != nil {
		return err
	}

	logger.Info("delete file %s in %s/%s successfully", args.filename, args.username, args.foldername)

	return nil
}
