package file

import "github.com/urfave/cli/v2"

// BeforeRegister is the command before the create command
func BeforeCreate(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	//foldername := c.Args().Get(2)
	//description := c.Args().Get(3) // optional
	return nil
}

// ActionRegister is the action to create a new file
func ActionCreate(c *cli.Context) error {
	// FIXME:
	return nil
}
