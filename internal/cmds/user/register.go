package user

import "github.com/urfave/cli/v2"

// BeforeRegister is the command before the register command
func BeforeRegister(c *cli.Context) error {
	// FIXME: Add validation
	//username := c.Args().Get(1)
	return nil
}

// ActionRegister is the action to register a new user
func ActionRegister(c *cli.Context) error {
	// FIXME: Add the register command
	return nil
}
