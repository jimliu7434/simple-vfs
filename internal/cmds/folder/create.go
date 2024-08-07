package folder

import "github.com/urfave/cli/v2"

// BeforeCreate is the command before the Create command
func BeforeCreate(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	//foldername := c.Args().Get(2)
	//description := c.Args().Get(3) // optional
	return nil
}

// ActionCreate is the action to create a new folder
func ActionCreate(c *cli.Context) error {
	// FIXME:
	return nil
}
