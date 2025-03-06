package service_test

import (
	"testing"

	"github.com/Anand078/moneytransfer/internal/model"
	"github.com/Anand078/moneytransfer/internal/service"
	"github.com/Anand078/moneytransfer/pkg/database"
	"go.uber.org/zap"
)

func TestExecuteTransfer(t *testing.T) {
	logger := zap.NewNop()
	store := database.NewStore(logger, []model.Account{
		{Name: "Mark", Balance: 100},
		{Name: "Jane", Balance: 50},
	})
	svc := service.NewTransferService(store, logger)

	req := model.TransferRequest{
		From:   "Mark",
		To:     "Jane",
		Amount: 30,
	}

	resp, err := svc.ExecuteTransfer(req)
	if err != nil || !resp.Success {
		t.Fatalf("Expected success but got error: %v", err)
	}

	accounts := store.GetAllAccounts()
	if accounts["Mark"].Balance != 70 {
		t.Errorf("Expected Mark's balance to be 70, got %.2f", accounts["Mark"].Balance)
	}
	if accounts["Jane"].Balance != 80 {
		t.Errorf("Expected Jane's balance to be 80, got %.2f", accounts["Jane"].Balance)
	}
}
