package folder

import "github.com/urfave/cli/v2"

// BeforeRename is the command before the rename command
func BeforeRename(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	//foldername := c.Args().Get(2)
	//newfoldername := c.Args().Get(3)
	return nil
}

// ActionRename is the action to rename a folder
func ActionRename(c *cli.Context) error {
	// FIXME:
	return nil
}
