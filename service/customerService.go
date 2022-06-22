package service

import (
	"github.com/kaliayev-proton/banking-go-hex/domain"
	"github.com/kaliayev-proton/banking-go-hex/errors"
	"github.com/kaliayev-proton/banking-go-hex/dto"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}