package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ssh-manager/internal/ssh"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an SSH alias",
	Long:  `Remove an SSH alias from your SSH config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide an alias to remove.")
			os.Exit(1)
		}
		alias := args[0]

		if err := ssh.RemoveAlias(alias); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Removed alias '%s' from SSH config file.\n", alias)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
