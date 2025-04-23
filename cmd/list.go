package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ssh-manager/internal/ssh"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all SSH aliases",
	Long:  `List all SSH aliases in your SSH config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := ssh.ListAliases()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(aliases) == 0 {
			fmt.Println("No SSH aliases found.")
			return
		}

		fmt.Println("SSH Aliases:")
		for _, alias := range aliases {
			fmt.Println("-", alias)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
