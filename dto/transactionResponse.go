package dto

type TransactionResponse struct {
	TransactionId string `json:"transaction_id"`
	AccountId string `json:"account_id"`
	Amount string `json:"amount"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
	CustomerId string `json:"-"`
}