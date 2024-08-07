// Package file is the package that contains the commands in category "file".
package file

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
	filename    string
	description string // optional
}

func (args *createArgs) IsValid() error {
	// if !Storage.IsValidFoldername(args.foldername) {
	// 	return fmt.Errorf("Invalid foldername %s", args.foldername)
	// }

	if !Storage.IsValidFilename(args.filename) {
		return fmt.Errorf("filename %s contains invalid chars", args.foldername)
	}

	return nil
}

// BeforeCreate is the command before the create command
func BeforeCreate(c *cli.Context) error {
	args := &createArgs{
		username:    c.Args().Get(0),
		foldername:  c.Args().Get(1),
		filename:    c.Args().Get(2),
		description: c.Args().Get(3), // optional
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)

	return nil
}

// ActionCreate is the action to create a new file
func ActionCreate(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*createArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	user, err := storage.GetUser(args.username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(args.foldername)
	if err != nil {
		return err
	}

	if err := folder.CreateFile(args.filename, args.description); err != nil {
		return err
	}

	util.Info("create file %s in %s/%s successfully\n", args.filename, args.username, args.foldername)

	return nil
}
