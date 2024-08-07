package file

import "github.com/urfave/cli/v2"

// BeforeList is the command before the List command
func BeforeList(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	return nil
}

// ActionList is the action to list all files in a folder
func ActionList(c *cli.Context) error {
	// FIXME:
	return nil
}
