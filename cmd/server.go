package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"yourapp/internal/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"yourapp/pkg/logger"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server.go",
	Long:  `Start the HTTP server.go that provides user management and authentication APIs.`,
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	// Initialize log
	log := logger.GetLogger()

	sv := server.NewServer()

	// Start server.go
	go func() {
		err := sv.Start()
		if err != nil {
			log.Fatal("Failed to start server.go", zap.Error(err))
		}
	}()

	// Wait for the interrupt signal to gracefully shut down the server.go
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server.go...")

	// Graceful shutdown
	if err := sv.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exiting")
}
