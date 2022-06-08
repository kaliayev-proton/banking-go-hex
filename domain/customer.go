package domain

import (
	"github.com/kaliayev-proton/banking-go-hex/errors"
)

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateOfBirth string
	Status string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errors.AppError)
}
