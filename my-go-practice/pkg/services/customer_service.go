package services

import "my-go-practice/pkg/models"

type CustomerService interface {
	GetCustomer(id int) (*models.Customer, error)
	UpdateBalance(id int, amount float64) error
	IsEligibleForLoan(id int) bool
}
