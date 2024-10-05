package main

import (
	"log"
	"github.com/randytjioe/merchant-bank-api/routes"
	"net/http"
)

func main() {
    routes.RegisterRoutes()
    log.Println("Server running on port 5050")
    http.ListenAndServe(":5050", nil)
}
