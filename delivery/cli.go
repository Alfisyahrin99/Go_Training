package delivery

import (
	"database/sql"
	"fmt"
	"godb/config"
	"godb/repository"
	"godb/usecase"

	"github.com/jmoiron/sqlx"
	// "strings"
)

func Run() {
	config := config.NewConfig()
	db := config.DbConnect()
	// customerRepo := repository.NewCustomerRepository(db)
	// customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	StoreCli(db)
	// ProdukCli(db)

	// //INSERT DATA CUSTOMER
	// var customer = model.Customer{
	// 	Name:    "Jhon LBF",
	// 	Balance: 60000,
	// }

	// err := customerUseCase.RegisterCustomer(&customer)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Success Create New Customer")
	// }\

	// customers, err := customerUseCase.GetAllCustomer()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, customer := range customers {
	// 	fmt.Println(strings.Repeat("==", 20))
	// 	fmt.Println("ID :", customer.Id)
	// 	fmt.Println("NAME :", customer.Name)
	// 	fmt.Println("BALANCE :", customer.Balance)
	// }
}

func StoreCli(db *sqlx.DB) {
	storeRepo := repository.NewStoreRepository(db)
	storeUseCase := usecase.NewStoreUseCase(storeRepo)

	//GET ALL STORE
	stores, err := storeUseCase.GetAllStore()
	if err != nil {
		panic(err.Error())
	}

	for _, store := range stores {
		fmt.Println(store)
	}
}

func ProdukCli(db *sql.DB) {
	produkRepo := repository.NewProdukRepository(db)
	produkUseCase := usecase.NewProdukUseCase(produkRepo)

	//GET ALL STORE
	produks, err := produkUseCase.GetAllProduk()
	if err != nil {
		panic(err.Error())
	}

	for _, produk := range produks {
		fmt.Println(produk)
	}
}
