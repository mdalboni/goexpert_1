package handlers

import (
	"context"
	"encoding/json"
	"goexpert_server_1/internal/services"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type QuotationHandler interface {
	GetQuotation(w http.ResponseWriter, r *http.Request)
}

type quotationHandler struct {
	quotationService services.QuotationService
	moneyService     services.MoneyService
}

func NewQuotationHandler(quotationService services.QuotationService, moneyService services.MoneyService) QuotationHandler {
	return &quotationHandler{
		quotationService: quotationService,
		moneyService:     moneyService,
	}
}

func (qh *quotationHandler) GetQuotation(w http.ResponseWriter, r *http.Request) {
	quotationResponse, err := qh.moneyService.GetDolarRealQuotation()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	quotation := quotationResponse.ToQuotation()
	qh.quotationService.SaveQuotation(ctx, &quotation)

	err = json.NewEncoder(w).Encode(quotationResponse)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}

	log.Info("Quotation queried with success")
}

type errorResponse struct {
	Message string `json:"message"`
}
