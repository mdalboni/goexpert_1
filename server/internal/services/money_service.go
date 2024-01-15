package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goexpert_server_1/internal/models"
	"net/http"
	"net/url"
	"time"
)

var baseURL = map[string]string{
	"prod": "https://economia.awesomeapi.com.br/",
}

type QuotationResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (qr *QuotationResponse) ToQuotation() models.Quotation {
	return models.Quotation{
		Code:       qr.Code,
		Codein:     qr.Codein,
		Name:       qr.Name,
		High:       qr.High,
		Low:        qr.Low,
		VarBid:     qr.VarBid,
		PctChange:  qr.PctChange,
		Bid:        qr.Bid,
		Ask:        qr.Ask,
		Timestamp:  qr.Timestamp,
		CreateDate: qr.CreateDate,
	}
}

type MoneyService interface {
	GetDolarRealQuotation() (QuotationResponse, error)
}

type moneyService struct {
	url        string
	httpClient *http.Client
}

func NewMoneyService(httpClient *http.Client) MoneyService {
	return &moneyService{
		url:        baseURL["prod"], // TODO get from env
		httpClient: httpClient,
	}
}

func (ms *moneyService) GetDolarRealQuotation() (QuotationResponse, error) {
	path := "json/last/USD-BRL"
	ctx, cancel := context.WithTimeout(context.Background(), 20000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ms.url+path, nil)
	if err != nil {
		return QuotationResponse{}, err
	}

	resp, err := ms.httpClient.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Err == context.DeadlineExceeded {
			return QuotationResponse{}, fmt.Errorf("request timed out: %w", err)
		}
		return QuotationResponse{}, err
	}
	defer resp.Body.Close()

	var result map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return QuotationResponse{}, err
	}

	var quotationResponse QuotationResponse
	err = json.Unmarshal(result["USDBRL"], &quotationResponse)
	if err != nil {
		return QuotationResponse{}, err
	}

	return quotationResponse, nil
}
