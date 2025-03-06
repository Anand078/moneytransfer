package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Anand078/moneytransfer/internal/api"
	"github.com/Anand078/moneytransfer/internal/model"
	"github.com/Anand078/moneytransfer/internal/service"
	"github.com/Anand078/moneytransfer/pkg/database"
	"go.uber.org/zap"
)

func setupTestServer() *api.Server {
	logger := zap.NewNop()

	initialAccounts := []model.Account{
		{Name: "Mark", Balance: 100},
		{Name: "Jane", Balance: 50},
	}

	store := database.NewStore(logger, initialAccounts)
	transferService := service.NewTransferService(store, logger)

	return api.NewServer(transferService, logger)
}

func TestMoneyTransferFlow(t *testing.T) {
	server := setupTestServer()

	testServer := httptest.NewServer(server.Handler())
	defer testServer.Close()

	baseURL := testServer.URL

	// Step 1: Check initial balances
	resp, err := http.Get(baseURL + "/accounts")
	if err != nil {
		t.Fatalf("Failed to fetch accounts: %v", err)
	}
	defer resp.Body.Close()

	var initialAccounts map[string]model.Account
	readJSON(t, resp.Body, &initialAccounts)

	if initialAccounts["Mark"].Balance != 100 || initialAccounts["Jane"].Balance != 50 {
		t.Fatalf("Unexpected initial balances: %+v", initialAccounts)
	}

	// Step 2: Perform a transfer
	transferReq := model.TransferRequest{
		From:   "Mark",
		To:     "Jane",
		Amount: 30,
	}
	body, _ := json.Marshal(transferReq)

	resp, err = http.Post(baseURL+"/transfer", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to perform transfer: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.StatusCode)
	}

	// Step 3: Check balances after transfer
	resp, err = http.Get(baseURL + "/accounts")
	if err != nil {
		t.Fatalf("Failed to fetch accounts after transfer: %v", err)
	}
	defer resp.Body.Close()

	var updatedAccounts map[string]model.Account
	readJSON(t, resp.Body, &updatedAccounts)

	if updatedAccounts["Mark"].Balance != 70 {
		t.Errorf("Expected Mark's balance to be 70, got %.2f", updatedAccounts["Mark"].Balance)
	}
	if updatedAccounts["Jane"].Balance != 80 {
		t.Errorf("Expected Jane's balance to be 80, got %.2f", updatedAccounts["Jane"].Balance)
	}
}

// Utility to read JSON response
func readJSON(t *testing.T, body io.Reader, target interface{}) {
	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
}

func TestConcurrentTransfers(t *testing.T) {
	server := setupTestServer()

	testServer := httptest.NewServer(server.Handler())
	defer testServer.Close()

	baseURL := testServer.URL

	// Set up concurrent transfers
	const concurrentTransfers = 50
	done := make(chan struct{})

	for i := 0; i < concurrentTransfers; i++ {
		go func() {
			transferReq := model.TransferRequest{
				From:   "Mark",
				To:     "Jane",
				Amount: 1,
			}
			body, _ := json.Marshal(transferReq)

			resp, err := http.Post(baseURL+"/transfer", "application/json", bytes.NewReader(body))
			if err != nil {
				t.Errorf("Transfer request failed: %v", err)
				return
			}
			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected 200 OK for transfer, got %d", resp.StatusCode)
			}

			done <- struct{}{}
		}()
	}

	// Wait for all transfers to complete
	for i := 0; i < concurrentTransfers; i++ {
		<-done
	}

	// Verify final balances
	resp, err := http.Get(baseURL + "/accounts")
	if err != nil {
		t.Fatalf("Failed to fetch accounts after concurrent transfers: %v", err)
	}
	defer resp.Body.Close()

	var finalAccounts map[string]model.Account
	readJSON(t, resp.Body, &finalAccounts)

	expectedMarkBalance := 100 - (1 * concurrentTransfers)
	expectedJaneBalance := 50 + (1 * concurrentTransfers)

	if finalAccounts["Mark"].Balance != float64(expectedMarkBalance) {
		t.Errorf("Expected Mark's balance to be %d, got %.2f", expectedMarkBalance, finalAccounts["Mark"].Balance)
	}
	if finalAccounts["Jane"].Balance != float64(expectedJaneBalance) {
		t.Errorf("Expected Jane's balance to be %d, got %.2f", expectedJaneBalance, finalAccounts["Jane"].Balance)
	}
}
