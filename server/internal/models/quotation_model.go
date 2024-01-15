package models

type Quotation struct {
	Code       string `gorm:"column:code" json:"code"`
	Codein     string `gorm:"column:codein" json:"codein"`
	Name       string `gorm:"column:name" json:"name"`
	High       string `gorm:"column:high" json:"high"`
	Low        string `gorm:"column:low" json:"low"`
	VarBid     string `gorm:"column:var_bid" json:"varBid"`
	PctChange  string `gorm:"column:pct_change" json:"pctChange"`
	Bid        string `gorm:"column:bid" json:"bid"`
	Ask        string `gorm:"column:ask" json:"ask"`
	Timestamp  string `gorm:"column:timestamp" json:"timestamp"`
	CreateDate string `gorm:"column:create_date" json:"create_date"`
	ID         uint   `gorm:"primarykey" json:"-"`
}
