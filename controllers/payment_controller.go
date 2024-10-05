package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go" // Ensure you have this package installed
	"github.com/randytjioe/merchant-bank-api/models"
	"github.com/randytjioe/merchant-bank-api/services"
)

// Payment handles the payment request
func Payment(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "Authorization header is missing",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNotSupported
		}
		// Replace with your secret key
		return []byte("your_secret_key"), nil
	})

	if err != nil || !token.Valid {
		response := ApiResponse{
			Status:  "error",
			Message: "Invalid token",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Decode the payment request
	var paymentRequest models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate the payment request fields
	if paymentRequest.MerchantID == "" {
		response := ApiResponse{
			Status:  "error",
			Message: "MerchantID is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if paymentRequest.Amount <= 0 {
		response := ApiResponse{
			Status:  "error",
			Message: "Amount must be greater than 0",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Call the payment service
	message, err := services.Payment(paymentRequest.MerchantID, paymentRequest.Amount)
	if err != nil {
		response := ApiResponse{
			Status:  "error",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Send successful response
	response := ApiResponse{
		Status:  "success",
		Message: "Payment successful",
		Data:    message, // Assuming message contains relevant payment information
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
