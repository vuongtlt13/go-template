package main

import (
	"log"
	"os"
	"yourapp/cmd"

	"github.com/joho/godotenv"

	"github.com/spf13/cobra"
)

func init() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found or could not be loaded (this is fine if running in production)")
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "yourapp"}
	rootCmd.AddCommand(cmd.AdminCmd)
	rootCmd.AddCommand(cmd.MigrateCmd)
	rootCmd.AddCommand(cmd.UserCmd) // Optional: keep if you want user server support
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
