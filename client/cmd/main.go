package main

import (
	"fmt"
	"goexpert_client_1/internal/services"
	"net/http"

	"github.com/labstack/gommon/log"
)

func main() {
	httpClient := &http.Client{}
	quotationService := services.NewQuotationService(httpClient)
	response, err := quotationService.GetQuotation()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	services.WriteFile("cotacao.txt", fmt.Sprintf("Dólar:%s", response.Bid))
}
