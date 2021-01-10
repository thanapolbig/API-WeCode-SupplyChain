package dashboard

import (
	"time"
)

const (
	PurchaseOrder = 110
	GoodsReceipt  = 120
	Invoice       = 210
	Seller        = 1001
	Buyer         = 1003

	Draft           = 1
	PendingApproval = 3
	PendingSubmited = 7
	ForModify       = 9
	Cancelled       = 13
	Accepted        = 15
	Submited        = 11
)

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type inputDashboard struct {
	FilterDateFrom time.Time `json:"filter_date_from"`
	FilterDateTo   time.Time `json:"filter_date_to"`
	LoginCompanyID int       `json:"login_company_id"`
	FilterCurrency string    `json:"filter_currency"`
}

type outputDashboard struct {
	DocumentType   string  `json:"document_type"`
	DocumentStatus int     `json:"document_status"`
	CountDocument  int     `json:"count_document"`
	Total          float64 `json:"total"`
}

type CompanyInfo struct {
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
	ShortName   string `json:"short_name"`
	CompanyType int    `json:"company_type"`
	Branch      string `json:"branch"`
	PhoneNumber string `json:"phone_number"`
	District    string `json:"district"`
	SubDistrict string `json:"sub_district"`
	ZipCode     string `json:"zip_code"`
	Province    string `json:"province"`
}
