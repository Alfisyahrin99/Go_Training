package usecase

import (
	"errors"
	"godb/model"
	"godb/repository"
)

type CustomerUseCase interface {
	RegisterCustomer(newCustomer *model.Customer) error
	GetAllCustomer() ([]model.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterCustomer(newCustomer *model.Customer) error {
	if len(newCustomer.Name) < 3 || len(newCustomer.Name) > 20 {
		return errors.New("Nama Minimal 3 sampai 20 Karakter")
	} else if newCustomer.Balance < 50000 {
		return errors.New("Mnimal Saldo 50.000")
	} else {
		return c.customerRepo.Insert(newCustomer)
	}
}

func (c *customerUseCase) GetAllCustomer() ([]model.Customer, error) {
	return c.customerRepo.GetAll()
}

func NewCustomerUseCase(customeerRepostitory repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customeerRepostitory,
	}
}
