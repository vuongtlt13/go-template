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

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Start the user HTTP server",
	Run:   runUserServer,
}

func runUserServer(cmd *cobra.Command, args []string) {
	// Initialize logger
	log := logger.GetLogger()

	// Load configuration
	cfg := config.GetConfig()

	// Create user server
	srv := server.NewUserServer(cfg, log)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal("Failed to start user server", "error", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Info("Shutting down user server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to shutdown user server", "error", err)
	}
}
