package dto

import (
	"github.com/kaliayev-proton/banking-go-hex/errors"
	"strings"
)

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}

	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return errors.NewValidationError("Account type should be checking or saving")
	}
	return nil
}