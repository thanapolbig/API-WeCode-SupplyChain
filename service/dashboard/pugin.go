package dashboard

import (
	mssql "api-wecode-supplychain/database/mssql"
	"time"
)

func logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = mssql.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}

func buyerDashboard(request inputDashboard) (result []outputDashboard, err error) {
	var ZeroValueDate time.Time
	buyerData := mssql.DB.Select("document_type ,document_status, COUNT(document_status) as count_document, ROUND(SUM(grand_total),2) as total").
		Table("document_header AS dh"). //110 = PO , 120 = GR , 210 = INV
		Where("buyer_company_ID = ?", request.LoginCompanyID).
		Where("document_type = ? and document_status in (?,?,?,?,?,?)", PurchaseOrder, Draft, PendingApproval, PendingSubmited, ForModify, Cancelled, Accepted).
		Or("document_type = ? and document_status in (?,?)", GoodsReceipt, Draft, Submited).
		Or("document_type = ? and document_status in (?,?,?,?)", Invoice, PendingSubmited, ForModify, Cancelled, Accepted).
		Group("document_type,document_status").Order("document_type")

	if request.FilterCurrency != "" {
		buyerData = buyerData.Where("currency = ?", request.FilterCurrency)
	} else {
		buyerData = buyerData.Where("currency = 'THB'")
	}

	if request.FilterDateFrom != ZeroValueDate {
		buyerData = buyerData.Where("create_date >= ?", request.FilterDateFrom)
	}

	if request.FilterDateTo != ZeroValueDate {
		buyerData = buyerData.Where("create_date <= ?", request.FilterDateTo)
	}

	if err = buyerData.Find(&result).Error; err != nil {
		return
	}
	return
	//return
}

func getSellserDashboard(request inputDashboard) (result []outputDashboard, err error) {
	var ZeroValueDate time.Time
	data := mssql.DB.Select("document_type,document_status , COUNT(*) as count_document ,SUM(grand_total) as total").
		Table("document_header AS dh").
		Where("((document_type = '110' and document_status in ('3','5','9','11','13')) or (document_type = '210' and document_status in ('1','5','9','11','13'))) and seller_company_id = ?", request.LoginCompanyID).
		Group("document_type,document_status").Order("document_type")

	if request.FilterCurrency != "" {
		data = data.Where("currency = ?", request.FilterCurrency)
	} else {
		data = data.Where("currency = 'THB'")
	}

	if request.FilterDateFrom != ZeroValueDate {
		data = data.Where("create_date >= ?", request.FilterDateFrom)
	}

	if request.FilterDateTo != ZeroValueDate {
		data = data.Where("create_date <= ?", request.FilterDateTo)
	}

	if err = data.Find(&result).Error; err != nil {
		return
	}

	return
}

func getCompanyById(request inputDashboard) (result []CompanyInfo, err error) {

	if err = mssql.DB.Select("*").
		Table("company").
		Where("id = ?", request.LoginCompanyID).
		Find(&result).Error; err != nil {
		return
	}

	return
}
