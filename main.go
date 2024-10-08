// Package: main is the package that contains the main function and the prompt function.
package main

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"

	filecmds "simple-vfs/internal/cmds/file"
	foldercmds "simple-vfs/internal/cmds/folder"
	usercmds "simple-vfs/internal/cmds/user"
	Storage "simple-vfs/internal/entity/storage"
	"simple-vfs/internal/logger"
)

var templates *promptui.PromptTemplates

// ErrEmpty is returned when the input is empty
var ErrEmpty = errors.New("empty input")

func init() {
	templates = &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
}

func main() {
	storage := Storage.New()

	for {
		err := prompt(&storage)
		if err != nil && err != ErrEmpty {
			// check if error is ctrl+c
			if err.Error() == "^C" {
				logger.Info("Goodbye!")
				return
			}
			logger.Error(err.Error())
		}
	}
}

func prompt(storage *Storage.Storage) error {
	cmdPrompt := promptui.Prompt{
		Label:     "#",
		Templates: templates,
	}

	cmdStr, err := cmdPrompt.Run()
	if err != nil {
		return err
	}

	if cmdStr == "" {
		return ErrEmpty
	}

	cmds := strings.Split(cmdStr, " ")

	cliApp := &cli.App{
		Name:  "vfs",
		Usage: "a simple virtual file system",
		Commands: []*cli.Command{
			{
				Name:      "register",
				Usage:     "register a new user",
				Category:  "User",
				Args:      true,
				ArgsUsage: "[username]",
				Before:    usercmds.BeforeRegister,
				Action:    usercmds.ActionRegister,
			},
			{
				Name:      "create-folder",
				Usage:     "create a new folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername] [description?]",
				Before:    foldercmds.BeforeCreate,
				Action:    foldercmds.ActionCreate,
			},
			{
				Name:      "delete-folder",
				Usage:     "delete a folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername]",
				Before:    foldercmds.BeforeDelete,
				Action:    foldercmds.ActionDelete,
			},
			{
				Name:      "list-folders",
				Usage:     "list all folders",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [--sort-name|sort-created asc|desc]",
				Before:    foldercmds.BeforeList,
				Action:    foldercmds.ActionList,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "sort-name",
						Usage:       "sort folders by name, asc | desc",
						DefaultText: "asc",
					},
					&cli.StringFlag{
						Name:        "sort-created",
						Usage:       "sort folders by created date, asc | desc",
						DefaultText: "asc",
					},
				},
			},
			{
				Name:      "rename-folder",
				Usage:     "rename a folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername] [new-folder-name]",
				Before:    foldercmds.BeforeRename,
				Action:    foldercmds.ActionRename,
			},
			{
				Name:      "create-file",
				Usage:     "create a new file",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [filename] [description?]",
				Before:    filecmds.BeforeCreate,
				Action:    filecmds.ActionCreate,
			},
			{
				Name:      "delete-file",
				Usage:     "delete a file",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [filename]",
				Before:    filecmds.BeforeDelete,
				Action:    filecmds.ActionDelete,
			},
			{
				Name:      "list-files",
				Usage:     "list all files in a folder",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [--sort-name|sort-created asc|desc]",
				Before:    filecmds.BeforeList,
				Action:    filecmds.ActionList,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "sort-name",
						Usage:       "sort folders by name, asc | desc",
						DefaultText: "asc",
					},
					&cli.StringFlag{
						Name:        "sort-created",
						Usage:       "sort folders by created date, asc | desc",
						DefaultText: "asc",
					},
				},
			},
		},
		Metadata: map[string]any{
			"storage": storage,
		},
		CommandNotFound: func(c *cli.Context, cmd string) {
			logger.Info("Command %s not found, some helps?\n\n", cmd)
			c.App.Command("help").Run(c)
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			logger.Info("Usage error: %s", err.Error())
			return nil
		},
		ExitErrHandler: func(c *cli.Context, err error) {
			if err != nil {
				logger.Error("%s", err.Error())
			}
		},
	}

	return cliApp.Run(append([]string{"vfs"}, cmds...))
}

// TODO: add "find parent"
