package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yourapp/internal/server"
	"yourapp/pkg/config"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"yourapp/pkg/logger"
)

var AdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Start the admin HTTP server",
	Long:  `Start the admin HTTP server that provides admin APIs for user, role, and permission management.`,
	Run:   runAdminServer,
}

func runAdminServer(cmd *cobra.Command, args []string) {
	// Initialize log
	log := logger.GetLogger()

	// Get config
	cfg := config.GetConfig()

	// Create and start server
	sv := server.NewAdminServer(cfg)

	// Start server
	go func() {
		err := sv.Start()
		if err != nil {
			log.Fatal("Failed to start admin server", zap.Error(err))
		}
	}()

	// Wait for the interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down admin server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := sv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Admin server exiting")
}
