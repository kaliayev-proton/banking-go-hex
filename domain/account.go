package domain

import (
	"github.com/kaliayev-proton/banking-go-hex/errors"
	"github.com/kaliayev-proton/banking-go-hex/dto"
)

type Account struct {
	AccountId string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errors.AppError)
	FindBy(accountId string) (*Account, *errors.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}