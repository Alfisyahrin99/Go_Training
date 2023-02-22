package repository

import (
	"godb/model"
	"godb/utils"

	"github.com/jmoiron/sqlx"
)

// sebagai kontraknya interface ini akan dijadikan sebagai acuan oleh layar lain
type CustomerRepository interface {
	Insert(newCustomer *model.Customer) error
	GetAll() ([]model.Customer, error)
}

// object implementasi dari interfacenya
type customerRepository struct {
	db *sqlx.DB
}

// method receiver untuk insert data
func (c *customerRepository) Insert(newCustomer *model.Customer) error {
	// query := "INSERT INTO mst_customer (name,balance) values ($1,$2)"
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, newCustomer.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) GetAll() ([]model.Customer, error) {
	rows, err := c.db.Query(utils.SELECT_CUSTOMER_LIST)
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Balance)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

// method untuk memanggil/menggunakan customerRepository
func NewCustomerRepository(db *sqlx.DB) CustomerRepository {

	return &customerRepository{
		db: db,
	}
}
