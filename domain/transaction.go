package domain

import (
	"github.com/kaliayev-proton/banking-go-hex/dto"
	"github.com/kaliayev-proton/banking-go-hex/errors"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type Transaction struct {
	TransactionId string `json:"transaction_id"`
	AccountId string `json:"account_id"`
	Amount string `json:"amount"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
}


func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		Amount: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (r dto.TransactionRequest) Validate() *errors.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errors.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errors.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}
