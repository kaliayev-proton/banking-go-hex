package dto

const WITHDRAWAL = "withdrawal"

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

type TransactionRequest struct {
	AccountId string `json:"account_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
	CustomerId string `json:"-"`
}

