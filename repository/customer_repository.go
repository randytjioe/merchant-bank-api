package repository

import (
	"encoding/json"
	"io/ioutil"
	"github.com/randytjioe/merchant-bank-api/models"
	"os"
)

func ReadCustomers() ([]models.Customer, error) {
    file, err := os.Open("data/customers.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)

    var customers []models.Customer
    json.Unmarshal(byteValue, &customers)
    return customers, nil
}

func WriteCustomers(customers []models.Customer) error {
    data, err := json.MarshalIndent(customers, "", "  ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile("data/customers.json", data, 0644)
    return err
}

// Update saldo pelanggan di file JSON
func UpdateCustomerBalance(updatedCustomer *models.Customer) error {
    customers, err := ReadCustomers()
    if err != nil {
        return err
    }

    // Update saldo pelanggan yang sesuai
    for i, customer := range customers {
        if customer.ID == updatedCustomer.ID {
            customers[i].Balance = updatedCustomer.Balance
            break
        }
    }

    // Simpan perubahan ke file
    return WriteCustomers(customers)
}

// Fungsi untuk menambahkan history ke file JSON
func AddHistory(history models.History) error {
    var histories []models.History

    // Baca file history jika ada, jika tidak ada buat baru
    if _, err := os.Stat("data/history.json"); err == nil {
        file, err := os.Open("data/history.json")
        if err != nil {
            return err
        }
        defer file.Close()

        byteValue, _ := ioutil.ReadAll(file)
        json.Unmarshal(byteValue, &histories)
    }

    // Tambahkan history baru
    histories = append(histories, history)

    // Simpan ke file
    data, err := json.MarshalIndent(histories, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/history.json", data, 0644)
}

