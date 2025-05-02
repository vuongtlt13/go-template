package cmd

import (
	"fmt"
	"log"
	"yourapp/pkg/database"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		dbCon := database.GetDatabase()

		if err := database.RunMigrations(dbCon, "migrations"); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}

		fmt.Println("Migration completed successfully.")
	},
}
