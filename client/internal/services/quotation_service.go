package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var baseURL = map[string]string{
	"prod": "http://localhost:8080/",
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

type QuotationService interface {
	GetQuotation() (QuotationResponse, error)
}

type quotationService struct {
	url        string
	httpClient *http.Client
}

func NewQuotationService(httpClient *http.Client) QuotationService {
	return &quotationService{
		url:        baseURL["prod"], // TODO get from env
		httpClient: httpClient,
	}
}

func (ms *quotationService) GetQuotation() (QuotationResponse, error) {
	path := "cotacao"
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ms.url+path, nil)
	if err != nil {
		return QuotationResponse{}, err
	}

	resp, err := ms.httpClient.Do(req)
	if err != nil {
		return QuotationResponse{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var quotationResponse QuotationResponse
		err = json.NewDecoder(resp.Body).Decode(&quotationResponse)
		if err != nil {
			return QuotationResponse{}, err
		}
		return quotationResponse, nil
	default:
		return QuotationResponse{}, err
	}

}
