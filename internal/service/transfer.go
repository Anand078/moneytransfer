package service

import (
	"fmt"

	"github.com/Anand078/moneytransfer/internal/model"
	"github.com/Anand078/moneytransfer/pkg/database"
	"go.uber.org/zap"
)

type TransferService struct {
	store  *database.Store
	logger *zap.Logger
}

func NewTransferService(store *database.Store, logger *zap.Logger) *TransferService {
	return &TransferService{
		store:  store,
		logger: logger,
	}
}

// ExecuteTransfer handles the business logic for a money transfer between accounts
func (s *TransferService) ExecuteTransfer(req model.TransferRequest) (model.TransferResponse, error) {
	s.logger.Info("Executing transfer request",
		zap.String("from", req.From),
		zap.String("to", req.To),
		zap.Float64("amount", req.Amount),
	)

	fromBalance, toBalance, err := s.store.Transfer(req.From, req.To, req.Amount)

	if err != nil {
		s.logger.Warn("Transfer failed", zap.Error(err))
		return model.TransferResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return model.TransferResponse{
		Success:     true,
		Message:     fmt.Sprintf("Successfully transferred $%.2f from %s to %s", req.Amount, req.From, req.To),
		FromBalance: fromBalance,
		ToBalance:   toBalance,
	}, nil
}

// GetAllAccounts returns the full list of accounts with balances
func (s *TransferService) GetAllAccounts() map[string]model.Account {
	return s.store.GetAllAccounts()
}
