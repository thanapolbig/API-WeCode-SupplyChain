package documents

import (
	"time"
)

const (
	PurchaseOrder = 110
	GoodsReceipt  = 120
	Invoice       = 210
	Seller        = 1001
	Buyer         = 1003

	Draft                = 1
	PendingApproval      = 3
	ForModify            = 5
	PendingSubmited      = 7
	ForModifyAfterSubmit = 9
	Cancelled            = 13
	Accepted             = 15
	Submited             = 11
)

type documentDetailReq struct {
	DocumentNumber  string `json:"document_number"`
	SellerCompanyId int64  `json:"seller_company_id"`
	BuyerCompanyId  int64  `json:"buyer_company_id"`
	DocumentType    int64  `json:"document_type"`
}
type documentDetail struct {
	SequenceNumber	int64	`json:"sequence_number"`
	DetailId		int64	`json:"detail_id"`
	Sku				string	`json:"sku"`
	ProductName		string	`json:"product_name"`
	Quantity		int64	`json:"quantity"`
	Unit			string	`json:"unit"`
	PricePerUnit	float64	`json:"price_per_unit"`
	TotalPrice		float64	`json:"total_price"`
}

type documentHeader struct {
	DocumentId			int64		`json:"document_id"`
	DocumentNumber		string		`json:"document_number"`
	DocumentStatus		string		`json:"document_status"`
	CreateDate			time.Time	`json:"create_date"`
	ExpireDate			time.Time	`json:"expire_date"`
	DeliveryDate		time.Time	`json:"delivery_date"`
	Net					float64		`json:"net"`
	VatPercent			int64		`json:"vat_percent"`
	VatValue			float64		`json:"vat_value"`
	Discount			float64		`json:"discount"`
	TotalAfterDiscount	float64		`json:"total_after_discount"`
	GrandTotal			float64		`json:"grand_total"`
	Currency			string		`json:"currency"`
	DeliveryAddress		string		`json:"delivery_address"`
	CompanyNameSeller	string		`json:"company_name_seller"`
	ShortNameSeller		string		`json:"short_name_seller"`
	BranchSeller		string		`json:"branch_seller"`
	PhoneNumberSeller	string		`json:"phone_number_seller"`
	DistrictSeller		string		`json:"district_seller"`
	SubDistrictSeller	string		`json:"sub_district_seller"`
	ZipCodeSeller		string		`json:"zip_code_seller"`
	ProvinceSeller		string		`json:"province_seller"`
	CompanyNameBuyer	string		`json:"company_name_buyer"`
	ShortNameBuyer		string		`json:"short_name_buyer"`
	BranchBuyer			string		`json:"branch_buyer"`
	PhoneNumberBuyer	string		`json:"phone_number_buyer"`
	DistrictBuyer		string		`json:"district_buyer"`
	SubDistrictBuyer	string		`json:"sub_district_buyer"`
	ZipCodeBuyer		string		`json:"zip_code_buyer"`
	ProvinceBuyer		string		`json:"province_buyer"`
}
type documentDetailRes struct {
	Header	documentHeader    `json:"header"`
	Detail	[]documentDetail `json:"detail"`
}
type draftDocument struct {
	Header	draftDocumentHeader		`json:"header"`
	Detail	[]draftDocumentDetail	`json:"detail"`
}
type draftDocumentHeader struct {
	DocumentNumber		string		`json:"document_number"`
	CreateDate			time.Time	`json:"create_date"`
	ExpireDate			time.Time	`json:"expire_date"`
	DeliveryDate		time.Time	`json:"delivery_date"`
	SellerCompanyId		int64		`json:"seller_company_id"`
	BuyerCompanyId		int64		`json:"buyer_company_id"`
	Net					float64		`json:"net"`
	VatPercent			int64		`json:"vat_percent"`
	VatValue			float64		`json:"vat_value"`
	Discount			float64		`json:"discount"`
	TotalAfterDiscount	float64		`json:"total_after_discount"`
	GrandTotal			float64		`json:"grand_total"`
	Currency			string		`json:"currency"`
	DocumentStatus		int64		`json:"document_status"`
	DocumentType		int64		`json:"document_type"`
	DeliveryAddress		string		`json:"delivery_address"`
	RecordStatus		bool		`json:"record_status"`
}
type draftDocumentDetail struct {
	//DetailId			int64	`json:"detail_id"`
	Sku					string	`json:"sku"`
	ProductName			string	`json:"product_name"`
	Quantity			int64	`json:"quantity"`
	Unit				string	`json:"unit"`
	PricePerUnit		float64	`json:"price_per_unit"`
	TotalPrice			float64	`json:"total_price"`
	SequenceNumber		int64	`json:"sequence_number"`
	DocumentHeaderId	int64	`json:"document_header_id"`
}

type messageResponse struct {
	Status				int64	`json:"status"`
	MessageCode			string	`json:"message_code"`
	MessageDescription	string	`json:"message_description"`
}

type documentListReq struct {
	DocumentNumber    string   `json:"document_number"`
	DocumentStatus    []int64  `json:"document_status"`
	SupplierName      string   `json:"supplier_name"`
	Currency          []string `json:"currency"`
	CreateDateStart   string   `json:"create_date_start"`
	CreateDateEnd     string   `json:"create_date_end"`
	ExpireDateStart   string   `json:"expire_date_start"`
	ExpireDateEnd     string   `json:"expire_date_end"`
	DeliveryDateStart string   `json:"delivery_date_start"`
	DeliveryDateEnd   string   `json:"delivery_date_end"`
	DocumentType      int64    `json:"document_type"`
	OrderBy           string   `json:"order_by"`
}

type documentList struct {
	Sequence       int64     `json:"sequence"`
	DocumentNumber string    `json:"document_number"`
	SupplierName   string    `json:"supplier_name"`
	CreateDate     time.Time `json:"create_date"`
	ExpireDate     time.Time `json:"expire_date"`
	DeliveryDate   time.Time `json:"delivery_date"`
	Total          float64   `json:"total"`
	Currency       string    `json:"currency"`
	DocumentStatus int       `json:"document_status"`
}
type PoCount struct {
	Count         int64   `json:"count"`
	DocumentType  string  `json:"document_type"`
	SumGrandTotal float64 `json:"sum_grand_total"`
}

type documentListRes struct {
	PoCount      PoCount        `json:"po_count"`
	DocumentList []documentList `json:"document_list"`
}

type inputDocumentListInv struct {
	DocumentNumber  string    `json:"document_number"`
	Currency        []string  `json:"currency"`
	SellerCompanyID []int64   `json:"seller_company_id"`
	DocumentStatus  []int64   `json:"document_status"`
	CreateDate      time.Time `json:"create_date"`
	CreateDate2     time.Time `json:"create_date2"`
	ExpireDate      time.Time `json:"expire_date"`
	ExpireDate2     time.Time `json:"expire_date2"`
	DeliveryDate    time.Time `json:"delivery_date"`
	DeliveryDate2   time.Time `json:"delivery_date2"`
}
type documentListInv struct {
	Sequence       int64     `json:"sequence"`
	DocumentId     int64     `json:"document_id"`
	DocumentNumber string    `json:"document_number"`
	CompanyId      int64     `json:"company_id"`
	CompanyName    string    `json:"company_name"`
	CreateDate     time.Time `json:"create_date"`
	ExpireDate     time.Time `json:"expire_date"`
	DeliveryDate   time.Time `json:"delivery_date"`
	GrandTotal     float64   `json:"grand_total"`
	Currency       string    `json:"currency"`
	DocumentStatus string    `json:"document_status"`
}
type documentInvRes struct {
	Header []documentInvCount
	Detail []documentListInv
}
type documentInvCount struct {
	DocumentIdTotal int64   `json:"document_id"`
	SumGrandTotal   float64 `json:"grand_total"`
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

//input status to postman
type inputStatus struct {
	LoginCompId    int64  `json:"login_comp_id"`
	TargetCompId   int64  `json:"target_comp_id"`
	DocumentNumber string `json:"document_number"`
	Note           string `json:"note"`
	ReasonCancel   string `json:"reason_cancel"`
	ReasonReject   string `json:"reason_reject"`
}

//message output
type MessageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
