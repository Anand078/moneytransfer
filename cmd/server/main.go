package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anand078/moneytransfer/internal/api"
	"github.com/Anand078/moneytransfer/internal/service"
	"github.com/Anand078/moneytransfer/pkg/config"
	"github.com/Anand078/moneytransfer/pkg/database"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conf := config.LoadConfig()

	store := database.NewStore(logger, config.ConvertToAccounts(conf.InitialBalances))
	service := service.NewTransferService(store, logger)
	server := api.NewServer(service, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.Start(conf.Server.Port); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	logger.Info("Server started", zap.Int("port", conf.Server.Port))

	<-ctx.Done()

	logger.Info("Received shutdown signal, gracefully stopping server...")

	// Gracefully shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx)

	logger.Info("Server stopped gracefully")
}
