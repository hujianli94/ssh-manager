package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh-manager",
	Short: "A simple SSH alias manager",
	Long: `ssh-manager is a CLI tool that helps you manage your SSH aliases easily.
You can add, list, and remove aliases.`,
}

// Execute 执行 root 命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
