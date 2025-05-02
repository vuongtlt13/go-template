package main

import (
	"github.com/spf13/cobra"
	"yourapp/cmd"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "User management application with background jobs",
	Long: `A user management application that provides authentication, user management,
and background job processing using Asynq.`,
}

func main() {
	// Add subcommands
	rootCmd.AddCommand(cmd.ServerCmd)
	rootCmd.AddCommand(cmd.MigrateCmd)
	//rootCmd.AddCommand(cmd.JobCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
