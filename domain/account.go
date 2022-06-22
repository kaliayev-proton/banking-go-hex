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
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}