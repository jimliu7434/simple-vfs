package folder

import (
	"context"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/util"

	"github.com/urfave/cli/v2"
)

type deleteArgs struct {
	username   string
	foldername string
}

func (args *deleteArgs) IsValid() error {
	// if !Storage.IsValidFoldername(args.foldername) {
	// 	return fmt.Errorf("Invalid foldername %s", args.foldername)
	// }

	return nil
}

// BeforeDelete is the command before the Delete command
func BeforeDelete(c *cli.Context) error {
	args := &deleteArgs{
		username:   c.Args().Get(1),
		foldername: c.Args().Get(2),
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionDelete is the action to delete a folder
func ActionDelete(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*deleteArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	if err := user.DelFolder(args.foldername); err != nil {
		return err
	}

	util.Info("delete folder %s successfully", args.foldername)

	return nil
}
