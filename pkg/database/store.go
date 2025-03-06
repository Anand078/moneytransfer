package database

import (
	"errors"
	"sync"

	"github.com/Anand078/moneytransfer/internal/model"
	"go.uber.org/zap"
)

type Store struct {
	accounts map[string]*model.Account
	mu       sync.RWMutex
	logger   *zap.Logger
}

func NewStore(logger *zap.Logger, initialBalances []model.Account) *Store {
	accounts := make(map[string]*model.Account)

	for _, acc := range initialBalances {
		accCopy := acc
		accounts[acc.Name] = &accCopy
	}
	return &Store{accounts: accounts, logger: logger}
}

func (s *Store) Transfer(from, to string, amount float64) (float64, float64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fromAcc, ok := s.accounts[from]
	if !ok {
		return 0, 0, errors.New("sender account not found")
	}

	toAcc, ok := s.accounts[to]
	if !ok {
		return 0, 0, errors.New("receiver account not found")
	}

	if from == to {
		return 0, 0, errors.New("cannot transfer to the same account")
	}

	if amount <= 0 {
		return 0, 0, errors.New("amount must be positive")
	}

	if fromAcc.Balance < amount {
		return 0, 0, errors.New("insufficient funds")
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount

	s.logger.Info("Transfer sucessful", zap.String("from", from), zap.String("to", to), zap.Float64("amount", amount))

	return fromAcc.Balance, toAcc.Balance, nil
}

func (s *Store) GetAllAccounts() map[string]model.Account {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]model.Account)

	for name, acc := range s.accounts {
		result[name] = *acc
	}
	return result
}
