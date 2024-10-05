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
    // Here, you should implement your logout logic, e.g., updating the user's session or status
    // If something goes wrong, return an error
    // For example:
    
    // Assuming you are just logging out, and there are no issues:
    // return nil if successful, or return an error if there are issues.
    loggedInCustomer = nil
    return nil // Return nil if logout is successful
}