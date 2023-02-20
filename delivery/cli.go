package delivery

import (
	"fmt"
	"godb/config"
	"godb/repository"
	"godb/usecase"
	"strings"
)

func Run() {
	config := config.NewConfig()
	db := config.DbConnect()
	customerRepo := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)

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

	customers, err := customerUseCase.GetAllCustomer()
	if err != nil {
		panic(err)
	}
	for _, customer := range customers {
		fmt.Println(strings.Repeat("==", 20))
		fmt.Println("ID :", customer.Id)
		fmt.Println("NAME :", customer.Name)
		fmt.Println("BALANCE :", customer.Balance)
	}
}
