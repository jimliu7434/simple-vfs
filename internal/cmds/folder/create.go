// Package folder is the package that contains the commands in category "folder".
package folder

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/util"
)

type createArgs struct {
	username    string
	foldername  string
	description string // optional
}

func (args *createArgs) IsValid() error {
	if !Storage.IsValidFoldername(args.foldername) {
		return fmt.Errorf("foldername %s contains invalid chars", args.foldername)
	}

	return nil
}

// BeforeCreate is the command before the Create command
func BeforeCreate(c *cli.Context) error {
	args := &createArgs{
		username:    c.Args().Get(0),
		foldername:  c.Args().Get(1),
		description: c.Args().Get(2), // optional
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionCreate is the action to create a new folder
func ActionCreate(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*createArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	if err := user.CreateFolder(args.foldername, args.description); err != nil {
		return err
	}

	util.Info("create folder %s successfully", args.foldername)

	return nil
}
