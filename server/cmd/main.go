package main

import (
	"goexpert_server_1/internal/database"
	"goexpert_server_1/internal/handlers"
	"goexpert_server_1/internal/services"
	"net/http"

	"github.com/labstack/gommon/log"
)

func main() {

	log.Info("Starting Server features.....")

	// Database Instance
	log.Info("Database starting process.....")
	db := database.SetupDatabase()
	defer database.CloseDatabase(db)
	log.Info("Database started with success.....")

	// Services Instances
	log.Info("Services starting.....")
	quotationService := services.NewQuotationService(db)
	httpClient := &http.Client{}
	moneyService := services.NewMoneyService(httpClient)
	log.Info("Services started with success.....")

	quotationHandler := handlers.NewQuotationHandler(quotationService, moneyService)

	http.HandleFunc("/cotacao", quotationHandler.GetQuotation)

	log.Info("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
