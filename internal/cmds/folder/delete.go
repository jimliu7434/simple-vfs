package folder

import "github.com/urfave/cli/v2"

// BeforeDelete is the command before the Delete command
func BeforeDelete(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	//foldername := c.Args().Get(2)
	//filename := c.Args().Get(3)
	return nil
}

// ActionDelete is the action to delete a folder
func ActionDelete(c *cli.Context) error {
	// FIXME:
	return nil
}
