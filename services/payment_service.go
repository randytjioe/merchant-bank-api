package services

import (
	"errors"
	"github.com/randytjioe/merchant-bank-api/models"
	"github.com/randytjioe/merchant-bank-api/repository"
	"time"
)

func Payment(merchantID string, amount float64) (string, error) {
    customer := GetLoggedInCustomer()
    if customer == nil {
        return "", errors.New("customer is not logged in")
    }

    if amount <= 0 {
        return "", errors.New("amount must be greater than zero")
    }

    if customer.Balance < amount {
        return "", errors.New("insufficient balance")
    }

    customer.Balance -= amount

    err := repository.UpdateCustomerBalance(customer)
    if err != nil {
        return "", err
    }
    
    history := models.History{
        ID:         time.Now().String(),
        CustomerID: customer.ID,
        Action:     "Payment",
        Amount:     amount,
        Timestamp:  time.Now().Format(time.RFC3339),
    }

    err = repository.AddHistory(history)
    if err != nil {
        return "", err
    }

    return "Payment successful", nil
}
