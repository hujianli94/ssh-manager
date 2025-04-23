package remove

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"ssh-manager/internal/ssh"
)

// Command remove 命令的定义
var Command = &cli.Command{
	Name:    "remove",
	Aliases: []string{"r"},
	Usage:   "Remove an SSH alias",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "alias",
			Aliases:  []string{"n"},
			Usage:    "Alias name to remove",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		alias := c.String("alias")
		if err := ssh.RemoveAlias(alias); err != nil {
			return err
		}

		fmt.Printf("Removed alias '%s' from SSH config file.\n", alias)
		return nil
	},
}
