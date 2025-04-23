package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ssh-manager/internal/ssh"
)

var removeAlias string

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an SSH alias",
	Long:  `Remove an SSH alias from your SSH config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if removeAlias == "" {
			fmt.Println("Please provide an alias to remove using --alias flag.")
			os.Exit(1)
		}

		if err := ssh.RemoveAlias(removeAlias); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Removed alias '%s' from SSH config file.\n", removeAlias)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&removeAlias, "alias", "a", "", "Alias name to remove (required)")
	removeCmd.MarkFlagRequired("alias")
}
