package file

import "github.com/urfave/cli/v2"

// BeforeDelete is the command before the Delete command
func BeforeDelete(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	return nil
}

// ActionDelete is the action to delete a file
func ActionDelete(c *cli.Context) error {
	// FIXME:
	return nil
}
