package repository

import (
	"godb/model"
	"godb/utils"

	"github.com/jmoiron/sqlx"
)

// Sebagai kontraknya, interface ini akan dijadikan sebagai acuan oleh layer lain
type CustomerRepository interface {
	Insert(newCustomer *model.Customer) error
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	UpdateById(customer model.Customer) error
	DeleteById(id string) error
}

// object implementasi dari interfacenya
type customerRepository struct {
	db *sqlx.DB
}

// method receiver untuk insert data
func (c *customerRepository) Insert(newCustomer *model.Customer) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, newCustomer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := c.db.Select(&customers, utils.SELECT_CUSTOMER_LIST)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) GetById(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow(utils.SELECT_CUSTOMER_ID, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Balance,
	)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) UpdateById(customer model.Customer) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) DeleteById(id string) error {
	_, err := c.db.Exec(utils.DELETE_CUSTOMER, id)
	if err != nil {
		return err
	}
	return nil
}

// method untuk memanggil/menggunakan customerRepository
func NewCustomerRepository(db *sqlx.DB) CustomerRepository {

	// repo := new(customerRepository)
	// repo.db = db

	// repo := &customerRepository{
	// 	db: db,
	// }

	return &customerRepository{
		db: db,
	}
}
