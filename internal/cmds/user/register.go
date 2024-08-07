// Package user is the package that contains the commands in category "user".
package user

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"
)

type createArgs struct {
	username string
}

func (args *createArgs) IsValid() error {
	if !Storage.IsValidUsername(args.username) {
		return fmt.Errorf("username %s contains invalid chars", args.username)
	}

	return nil
}

// BeforeRegister is the command before the register command
func BeforeRegister(c *cli.Context) error {
	args := &createArgs{
		username: c.Args().Get(0),
	}

	if err := args.IsValid(); err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, argsKey, args)
	return nil
}

// ActionRegister is the action to register a new user
func ActionRegister(c *cli.Context) error {
	args := c.Context.Value(argsKey).(*createArgs)
	storage := c.App.Metadata["storage"].(*Storage.Storage)

	if user, _ := storage.GetUser(args.username); user != nil {
		return fmt.Errorf("user %s has already existed", args.username)
	}

	if err := storage.CreateUser(args.username); err != nil {
		return err
	}

	logger.Info("add user %s successfully", args.username)
	return nil
}
