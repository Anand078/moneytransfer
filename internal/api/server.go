package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Anand078/moneytransfer/internal/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
)

type Server struct {
	handlers *Handlers
	logger   *zap.Logger
	server   *http.Server
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
}

func NewServer(service *service.TransferService, logger *zap.Logger) *Server {
	handlers := NewHandlers(service, logger)

	mux := http.NewServeMux()

	// Register endpoints
	mux.HandleFunc("/accounts", handlers.GetAllAccounts)
	mux.HandleFunc("/transfer", handlers.TransferMoney)

	// Observability endpoints
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return &Server{
		handlers: handlers,
		logger:   logger,
		server: &http.Server{
			Handler: mux,
		},
	}
}

// Start runs the server.
func (s *Server) Start(port int) error {
	s.server.Addr = fmt.Sprintf(":%d", port)
	s.logger.Info("Starting server", zap.String("address", s.server.Addr))
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("Server shutdown failed", zap.Error(err))
	} else {
		s.logger.Info("Server gracefully stopped")
	}
}
