package services

import (
	"context"
	"goexpert_server_1/internal/models"

	"gorm.io/gorm"
)

type QuotationService interface {
	SaveQuotation(ctx context.Context, quotation *models.Quotation) error
	GetQuotation(ctx context.Context, id uint) (*models.Quotation, error)
}

type quotationService struct {
	db *gorm.DB
}

func NewQuotationService(db *gorm.DB) QuotationService {
	return &quotationService{db: db}
}

func (qs *quotationService) SaveQuotation(ctx context.Context, quotation *models.Quotation) error {
	return qs.db.WithContext(ctx).Create(quotation).Error
}

func (qs *quotationService) GetQuotation(ctx context.Context, id uint) (*models.Quotation, error) {
	var quotation models.Quotation
	err := qs.db.WithContext(ctx).First(&quotation, id).Error
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}
