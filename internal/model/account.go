package model

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type TransferRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type TransferResponse struct {
	Success     bool    `json:"success"`
	Message     string  `json:"message"`
	FromBalance float64 `json:"from_balance,omitempty"`
	ToBalance   float64 `json:"to_balance,omitempty"`
}
