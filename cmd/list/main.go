package list

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"ssh-manager/internal/ssh"
)

// Command list 命令的定义
var Command = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List all SSH aliases",
	Action: func(c *cli.Context) error {
		aliases, err := ssh.ListAliases()
		if err != nil {
			return err
		}

		if len(aliases) == 0 {
			fmt.Println("No SSH aliases found.")
			return nil
		}

		fmt.Println("SSH Aliases:")
		for _, alias := range aliases {
			fmt.Println("-", alias)
		}

		return nil
	},
}
