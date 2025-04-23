package add

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"ssh-manager/internal/config"
	"ssh-manager/internal/ssh"
)

// Command add 命令的定义
var Command = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add a new SSH alias",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "alias",
			Aliases:  []string{"n"},
			Usage:    "Alias name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "host",
			Aliases:  []string{"H"},
			Usage:    "Host address",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Username",
			Value:   os.Getenv("USER"), // 默认使用当前用户名
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Port number",
			Value:   22, // 默认端口号为 22
		},
	},
	Action: func(c *cli.Context) error {
		alias := c.String("alias")
		host := c.String("host")
		user := c.String("user")
		port := c.Int("port")

		sshConfig := config.SSHConfig{
			Host:     alias,
			HostName: host,
			User:     user,
			Port:     port,
		}

		if err := ssh.AddAlias(sshConfig); err != nil {
			return err
		}

		fmt.Printf("Added alias '%s' to SSH config file.\n", alias)
		return nil
	},
}
