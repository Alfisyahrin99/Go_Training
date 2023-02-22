package delivery

import (
	"database/sql"
	"fmt"
	"godb/config"
	"godb/repository"
	"godb/usecase"
	"strings"

	"github.com/jmoiron/sqlx"
	// "strings"
)

func Run() {
	config := config.NewConfig()
	db := config.DbConnect()

	// StoreCli(db)
	CustomerCli(db)

}

func StoreCli(db *sql.DB) {
	productRepo := repository.NewProductRepository(db)
	productUc := usecase.NewProductUseCase(productRepo)

	storeRepo := repository.NewStoreRepository(db)
	storeUseCase := usecase.NewStoreUseCase(storeRepo, productUc)

	// GET ALL STORE
	storeProducts, err := storeUseCase.GetAllStoreWithProducts()
	if err != nil {
		panic(err.Error())
	}

	for _, store := range storeProducts {
		fmt.Println("Store Id :", store.Id)
		fmt.Println("Store Name :", store.StoreName)
		for _, product := range store.Products {
			println()
			fmt.Println("  Product Name :", product.Name)
			fmt.Println("  Product price :", product.Price)
			fmt.Println("  Product Stock :", product.Stock)
		}
		println()
	}
}

func ProductCli(db *sql.DB) {
	productRepo := repository.NewProductRepository(db)
	ProductUc := usecase.NewProductUseCase(productRepo)

	products, err := ProductUc.GetAllProduct()
	if err != nil {
		panic(err.Error())
	}

	for _, product := range products {
		fmt.Println(product)
	}
}

func CustomerCli(db *sqlx.DB) {
	customerRepo := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)

	// INSERT DATA CUSTOMER
	// var customer = model.Customer {
	// 	Name: "Jhon LBF",
	// 	Balance: 60000,
	// }
	// err := customerUseCase.RegisterCustomer(&customer)

	// if err != nil {
	// 	panic(err)
	// }else{
	// 	fmt.Println("Success Create new Customer")
	// }

	//GET LIST CUSTOMER
	customers, err := customerUseCase.GetAllCustomer()
	if err != nil {
		panic(err)
	}
	for _, customer := range customers {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Println("ID :", customer.Id)
		fmt.Println("NAME :", customer.Name)
		fmt.Println("BALANCE :", customer.Balance)
	}

	//GET CUSTOMER BY ID
	// customer,err := customerUseCase.GetCustomerById("13")
	// if err != nil {
	// 	panic(err)
	// }else{
	// 	fmt.Println(customer)
	// }

	//UPDATE CUSTOMER BY ID
	// customerUpdated := model.Customer{
	// 	Id:      "16",
	// 	Name:    "Robert",
	// 	Balance: 500000,
	// }
	// err := customerUseCase.UpdateCustomerById(customerUpdated)
	// if err != nil {
	// 	panic(err.Error())
	// }

	//DELETE CUSTOMER BY ID
	// err := customerUseCase.DeleteCustomerById("13")
	// if err != nil {
	// 	panic(err.Error())
	// }else{
	// 	fmt.Println("Success Delete Customer")
	// }
}
