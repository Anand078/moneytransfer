package api

import (
	"encoding/json"
	"net/http"

	"github.com/Anand078/moneytransfer/internal/model"
	"github.com/Anand078/moneytransfer/internal/service"
	"go.uber.org/zap"
)

type Handlers struct {
	service *service.TransferService
	logger  *zap.Logger
}

func NewHandlers(service *service.TransferService, logger *zap.Logger) *Handlers {
	return &Handlers{service: service, logger: logger}
}

// GetAccounts retrieves all accounts and their balances.
func (h *Handlers) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := h.service.GetAllAccounts()
	h.respondJSON(w, http.StatusOK, accounts)
}

// TransferMoney handles money transfers between accounts.
func (h *Handlers) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var req model.TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.From == req.To {
		h.respondError(w, http.StatusBadRequest, "cannot transfer to the same account")
		return
	}

	resp, err := h.service.ExecuteTransfer(req)
	if err != nil {
		h.respondJSON(w, http.StatusBadRequest, resp)
		return
	}

	h.respondJSON(w, http.StatusOK, resp)
}

// respondJSON sends a JSON response
func (h *Handlers) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError sends a structured error response
func (h *Handlers) respondError(w http.ResponseWriter, status int, message string) {
	h.logger.Warn("Request failed", zap.Int("status", status), zap.String("error", message))
	h.respondJSON(w, status, map[string]string{"error": message})
}
