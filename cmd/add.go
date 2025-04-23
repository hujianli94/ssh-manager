package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ssh-manager/internal/config"
	"ssh-manager/internal/ssh"
)

var (
	alias string
	host  string
	user  string
	port  int
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new SSH alias",
	Long:  `Add a new SSH alias to your SSH config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		sshConfig := config.SSHConfig{
			Host:     alias,
			HostName: host,
			User:     user,
			Port:     port,
		}

		if err := ssh.AddAlias(sshConfig); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Added alias '%s' to SSH config file.\n", alias)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 定义 flags
	addCmd.Flags().StringVarP(&alias, "alias", "n", "", "Alias name (required)")
	addCmd.Flags().StringVarP(&host, "host", "H", "", "Host address (required)")
	addCmd.Flags().StringVarP(&user, "user", "u", os.Getenv("USER"), "Username (default: current user)")
	addCmd.Flags().IntVarP(&port, "port", "p", 22, "Port number (default: 22)")

	// 标记 required flags
	addCmd.MarkFlagRequired("alias")
	addCmd.MarkFlagRequired("host")
}
