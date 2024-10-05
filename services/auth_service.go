package services

import (
	"errors"

	"github.com/randytjioe/merchant-bank-api/models"
	"github.com/randytjioe/merchant-bank-api/repository"
)

var loggedInCustomer *models.Customer

func Login(username, password string) (*models.Customer, error) {
    customers, err := repository.ReadCustomers()
    if err != nil {
        return nil, err
    }

    for _, customer := range customers {
        if customer.Username == username && customer.Password == password {
            loggedInCustomer = &customer
            return &customer, nil
        }
    }
    return nil, errors.New("invalid credentials")
}


func GetLoggedInCustomer() *models.Customer {
    return loggedInCustomer
}
func Logout() error {
    loggedInCustomer = nil
    return nil 
}