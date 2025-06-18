package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"yourapp/internal/server"
	"yourapp/pkg/config"
	"yourapp/pkg/logger"
)

var AdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Start the admin HTTP server",
	Run:   runAdminServer,
}

func runAdminServer(cmd *cobra.Command, args []string) {
	// Initialize logger
	log := logger.GetLogger()

	// Load configuration
	cfg := config.GetConfig()

	// Create admin server
	srv := server.NewAdminServer(cfg, log)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal("Failed to start admin server", "error", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Info("Shutting down admin server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to shutdown admin server", "error", err)
	}
}
