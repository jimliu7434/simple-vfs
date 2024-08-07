package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"

	filecmds "simple-vfs/internal/cmds/file"
	foldercmds "simple-vfs/internal/cmds/folder"
	usercmds "simple-vfs/internal/cmds/user"
	Storage "simple-vfs/internal/entity/storage"
)

var templates *promptui.PromptTemplates

// ErrEmpty is returned when the input is empty
var ErrEmpty = fmt.Errorf("empty input")

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
		if err := prompt(&storage); err != nil && err != ErrEmpty {
			fmt.Println(err.Error())
		}
	}
}

func prompt(storage *Storage.Storage) error {
	cmdPrompt := promptui.Prompt{
		Label:     "context file location",
		Templates: templates,
	}

	cmdStr, err := cmdPrompt.Run()
	if err != nil {
		return err
	}

	if cmdStr == "" {
		return ErrEmpty
	}

	cmd := strings.Split(cmdStr, " ")

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
	}

	return cliApp.Run(cmd)
}
