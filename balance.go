package neero

// BalanceResponse represents the account balance information.
type BalanceResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
