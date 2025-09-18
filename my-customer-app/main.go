package main

import (
	"my-customer-app/pkg/services"
)

func main() {
	// Initialize customer service
	customerService := services.NewCustomerService()

	// Simulate operations
	customerService.PrintCustomer(1)
	customerService.UpdateBalance(1, 500)
	customerService.PrintCustomer(1)

	customerService.PrintCustomer(2)
	customerService.UpdateBalance(2, -200)
	customerService.PrintCustomer(2)

	customerService.PrintCustomer(3)
	customerService.UpdateBalance(3, 1000)
	customerService.PrintCustomer(3)
}
