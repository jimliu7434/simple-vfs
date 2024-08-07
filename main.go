package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
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
	for {
		if err := prompt(); err != nil && err != ErrEmpty {
			fmt.Println(err.Error())
		}
	}
}

func prompt() error {
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

	// TODO: Add the command to the cli app
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
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the register command
					return nil
				},
			},
			{
				Name:      "create-folder",
				Usage:     "create a new folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername] [description?]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					//foldername := c.Args().Get(2)
					//description := c.Args().Get(3) // optional
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the create-folder command
					return nil
				},
			},
			{
				Name:      "delete-folder",
				Usage:     "delete a folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					//foldername := c.Args().Get(2)
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the delete-folder command
					return nil
				},
			},
			{
				Name:      "list-folders",
				Usage:     "list all folders",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [--sort-name|sort-created asc|desc]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)

					return nil
				},
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
				Action: func(c *cli.Context) error {
					// FIXME: Add the list-folders command
					return nil
				},
			},
			{
				Name:      "rename-folder",
				Usage:     "rename a folder",
				Category:  "Folder",
				Args:      true,
				ArgsUsage: "[username] [foldername] [new-folder-name]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					//foldername := c.Args().Get(2)
					//newfoldername := c.Args().Get(3)
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the rename-folder command
					return nil
				},
			},
			{
				Name:      "create-file",
				Usage:     "create a new file",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [filename] [description?]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					//foldername := c.Args().Get(2)
					//description := c.Args().Get(3) // optional
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the create-file command
					return nil
				},
			},
			{
				Name:      "delete-file",
				Usage:     "delete a file",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [filename]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					//username := c.Args().Get(1)
					//foldername := c.Args().Get(2)
					//filename := c.Args().Get(3)
					return nil
				},
				Action: func(c *cli.Context) error {
					// FIXME: Add the delete-file command
					return nil
				},
			},
			{
				Name:      "list-files",
				Usage:     "list all files",
				Category:  "File",
				Args:      true,
				ArgsUsage: "[username] [foldername] [--sort-name|sort-created asc|desc]",
				Before: func(c *cli.Context) error {
					// FIXME: Add validation
					// username := c.Args().Get(1)
					// foldername := c.Args().Get(2)

					return nil
				},
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
				Action: func(c *cli.Context) error {
					// FIXME: Add the list-files command
					return nil
				},
			},
		},
	}

	return cliApp.Run(cmd)
}
