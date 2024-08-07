package folder

import (
	"context"
	"fmt"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"

	"github.com/urfave/cli/v2"
)

type renameArgs struct {
	username      string
	foldername    string
	newfoldername string
}

func (args *renameArgs) IsValid() error {
	// if !Storage.IsValidFoldername(args.foldername) {
	// 	return fmt.Errorf("Invalid foldername %s", args.foldername)
	// }

	if !Storage.IsValidFoldername(args.newfoldername) {
		return fmt.Errorf("foldername %s contains invalid chars", args.newfoldername)
	}

	return nil
}

// BeforeRename is the command before the rename command
func BeforeRename(c *cli.Context) error {
	args := &renameArgs{
		username:      c.Args().Get(0),
		foldername:    c.Args().Get(1),
		newfoldername: c.Args().Get(2),
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionRename is the action to rename a folder
func ActionRename(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*renameArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	if err := user.RenameFolder(args.foldername, args.newfoldername); err != nil {
		return err
	}

	logger.Info("rename %s to %s successfully", args.foldername, args.newfoldername)

	return nil
}
