package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yourapp/internal/server"
	"yourapp/pkg/config"
	"yourapp/pkg/i18n"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"yourapp/pkg/logger"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Start the user HTTP server",
	Long:  `Start the user HTTP server that provides user-facing APIs.`,
	Run:   runUserServer,
}

func runUserServer(cmd *cobra.Command, args []string) {
	// Initialize log
	log := logger.GetLogger()

	// Get config
	cfg := config.GetConfig()

	// Initialize i18n
	i18n.Init(&cfg.I18n)

	// Create and start server
	sv := server.NewUserServer(cfg)

	// Start server
	go func() {
		err := sv.Start()
		if err != nil {
			log.Fatal("Failed to start user server", zap.Error(err))
		}
	}()

	// Wait for the interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down user server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := sv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("User server exiting")
}
