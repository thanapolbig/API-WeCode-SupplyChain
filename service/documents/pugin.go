package documents

import (
	mssql "api-wecode-supplychain/database/mssql"
	"time"
	"net/http"
)

func getDocDetail(request documentDetailReq) (result []documentDetail, err error) {

	if err = mssql.DB.Select("dd.sequence_number, dd.id AS detail_id, dd.sku, dd.product_name, dd.quantity, dd.unit, dd.price_per_unit, dd.total_price").
		Table("malar.dbo.document_detail AS dd").
		Joins("LEFT JOIN malar.dbo.document_header AS dh ON dd.document_header_id  = dh.document_id").
		Where("dh.document_number = ? AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? AND dh.record_status =1", request.DocumentNumber, request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).
		Find(&result).Error; err != nil {
		return
	}

	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
	//	return
	//}

	return
}

func getDocHeader(request documentDetailReq, status []int) (result documentHeader, err error) {
	data := mssql.DB.Select("dh.document_id, dh.document_number, c.description AS document_status, dh.create_date, dh.expire_date, dh.delivery_date, dh.net, dh.vat_percent, dh.vat_value, dh.discount, dh.grand_total, dh.currency, dh.delivery_address, cs.company_name AS company_name_seller, cs.short_name AS short_name_seller, cs.branch AS branch_seller, cs.phone_number AS phone_number_seller, cs.district AS district_seller, cs.sub_district AS sub_district_seller, cs.zip_code AS zip_code_seller, cs.province AS province_seller, cb.company_name AS company_name_buyer, cb.short_name AS short_name_buyer, cb.branch AS branch_buyer, cb.phone_number AS phone_number_buyer, cb.district AS district_buyer, cb.sub_district AS sub_district_buyer, cb.zip_code AS zip_code_buyer, cb.province AS province_buyer").
		Table("malar.dbo.document_header AS dh").
		Joins("LEFT JOIN malar.dbo.company AS cs on dh.seller_company_id = cs.id LEFT JOIN malar.dbo.company AS cb on dh.seller_company_id = cb.id LEFT JOIN malar.dbo.config AS c ON dh.document_status = c.status_code").
		Where("dh.document_number = ? AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? ", request.DocumentNumber, request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType)

	if len(status) != 0 || status != nil {
		data = data.Where("document_status IN (?)", status)
	}
	if err = data.Find(&result).Error; err != nil {
		return
	}

	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
	//	return
	//}
	return
}
func createDocumentHeader(request draftDocumentHeader) (result messageResponse, err error) {

	if err = mssql.DB.
		Table("malar.dbo.document_header").
		Save(&request).Error; err != nil {
		return
	}

	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
	//	return
	//}
	result.Status = http.StatusCreated
	result.MessageCode = "0000"
	result.MessageDescription = "create document Header success"

	return
}

func createDocumentDetail(request []draftDocumentDetail) (result messageResponse, err error) {
	for _, detail := range request {
		if err = mssql.DB.
			Table("malar.dbo.document_detail").
			Save(&detail).Error; err != nil {
			return
		}
	}
	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
	//	return
	//}
	result.Status = http.StatusCreated
	result.MessageCode = "0000"
	result.MessageDescription = "create document detail success"
	return
}

func checkDocumentHeader(request draftDocumentHeader) (result draftDocumentHeader, err error) {

	//if err = mssql.DB.Select("dh.document_id, dh.document_number, c.description AS document_status, dh.create_date, dh.expire_date, dh.delivery_date, dh.net, dh.vat_percent, dh.vat_value, dh.discount, dh.total_after_discount, dh.grand_total, dh.currency, dh.delivery_address, cs.company_name AS company_name_seller, cs.short_name AS short_name_seller, cs.branch AS branch_seller, cs.phone_number AS phone_number_seller, cs.district AS district_seller, cs.sub_district AS sub_district_seller, cs.zip_code AS zip_code_seller, cs.province AS province_seller, cb.company_name AS company_name_buyer, cb.short_name AS short_name_buyer, cb.branch AS branch_buyer, cb.phone_number AS phone_number_buyer, cb.district AS district_buyer, cb.sub_district AS sub_district_buyer, cb.zip_code AS zip_code_buyer, cb.province AS province_buyer").
	//	Table("document_header AS dh").
	//	Joins("LEFT JOIN company AS cs on dh.seller_company_id = cs.id LEFT JOIN company AS cb on dh.seller_company_id = cb.id LEFT JOIN config AS c ON dh.document_status = c.status_code").
	//	Where("dh.document_number LIKE ? AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? AND dh.record_status =1", request.DocumentNumber + "%", request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).
	//	//Order("dh.document_number desc").
	//	Last(&result).Error; err != nil {
	//	return
	//}
	if err = mssql.DB.Select("document_number, create_date, expire_date, delivery_date, seller_company_id, buyer_company_id, net, vat_percent, vat_value, discount, total_after_discount, grand_total, currency, document_status, document_type, delivery_address, record_status").
		Table("malar.dbo.document_header").
		//Joins("LEFT JOIN company AS cs on dh.seller_company_id = cs.id LEFT JOIN company AS cb on dh.seller_company_id = cb.id LEFT JOIN config AS c ON dh.document_status = c.status_code").
		Where("document_number LIKE ? AND seller_company_id = ? AND buyer_company_id = ? AND document_type = ? AND record_status =1", request.DocumentNumber+"%", request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).
		Order("document_number desc ").
		Last(&result).Error; err != nil {
		return
	}

	//if err = mssql.DB.Raw("SELECT dh.document_id, dh.document_number, c.description AS document_status, dh.create_date, dh.expire_date, dh.delivery_date, dh.net, dh.vat_percent, dh.vat_value, dh.discount, dh.total_after_discount, dh.grand_total, dh.currency, dh.delivery_address, cs.company_name AS company_name_seller, cs.short_name AS short_name_seller, cs.branch AS branch_seller, cs.phone_number AS phone_number_seller, cs.district AS district_seller, cs.sub_district AS sub_district_seller, cs.zip_code AS zip_code_seller, cs.province AS province_seller, cb.company_name AS company_name_buyer, cb.short_name AS short_name_buyer, cb.branch AS branch_buyer, cb.phone_number AS phone_number_buyer, cb.district AS district_buyer, cb.sub_district AS sub_district_buyer, cb.zip_code AS zip_code_buyer, cb.province AS province_buyer " +
	//	"FROM document_header AS dh LEFT JOIN company AS cs on dh.seller_company_id = cs.id LEFT JOIN company AS cb on dh.seller_company_id = cb.id LEFT JOIN config AS c ON dh.document_status = c.status_code" +
	//	"dh.document_number LIKE ?% AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? AND dh.record_status =1" +
	//	"ORDER BY dh.document_number ASC", request.DocumentNumber, request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).First(&result).Error; err != nil {
	//	return
	//}

	return
}

func getPODocList(request documentListReq) (result []documentList, err error) {

	if err = mssql.DB.Select("ROW_NUMBER() OVER(ORDER BY document_number) AS sequence,document_number,c.company_name AS supplier_name,create_date ,expire_date ,delivery_date ,grand_total AS total,currency ,document_status").
		Table("document_header AS dh,company AS c").
		Where("dh.document_type = ? AND currency IN ( ? ) AND company_name = ? AND document_status IN ( ? ) AND create_date BETWEEN ? AND ? AND expire_date BETWEEN ? AND ? AND delivery_date BETWEEN ? AND ?", request.DocumentType, request.Currency, request.SupplierName, request.DocumentStatus, request.CreateDateStart, request.CreateDateEnd, request.ExpireDateStart, request.ExpireDateEnd, request.DeliveryDateStart, request.DeliveryDateEnd).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func getPODocListCount(request documentListReq) (result PoCount, err error) {

	if err = mssql.DB.Select("COUNT(document_number) AS count,SUM(grand_total) AS sum_grand_total").
		Table("document_header AS dh,company AS c ").
		Where("dh.document_type = ? AND currency IN ( ? ) AND company_name = ? AND document_status IN ( ? ) AND create_date BETWEEN ? AND ? AND expire_date BETWEEN ? AND ? AND delivery_date BETWEEN ? AND ?", request.DocumentType, request.Currency, request.SupplierName, request.DocumentStatus, request.CreateDateStart, request.CreateDateEnd, request.ExpireDateStart, request.ExpireDateEnd, request.DeliveryDateStart, request.DeliveryDateEnd).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func documentListInvGet(request inputDocumentListInv) (result []documentListInv, err error) {

	if err = mssql.DB.Select("ROW_NUMBER() OVER(ORDER BY dh.document_id) AS sequence,dh.document_id, dh.document_number, c2.id AS company_id, c2.company_name, dh.create_date, dh.expire_date, dh.delivery_date, dh.grand_total, dh.currency, c.description AS document_status").
		Table("document_header AS dh, config AS c, company AS c2").
		Where("dh.document_type = 210 AND dh.document_number LIKE ? AND dh.currency IN ( ? ) AND dh.seller_company_id = ? AND dh.seller_company_id = c2.id AND dh.document_status IN ( ? ) AND c.status_code = dh.document_status AND dh.create_date BETWEEN ? and ? AND dh.expire_date BETWEEN ? and ? AND dh.delivery_date BETWEEN ? and ? ", "%"+request.DocumentNumber+"%", request.Currency, request.SellerCompanyID, request.DocumentStatus, request.CreateDate, request.CreateDate2, request.ExpireDate, request.ExpireDate2, request.DeliveryDate, request.DeliveryDate2).
		Order("dh.document_id").
		//Offset(0).
		Find(&result).Error; err != nil {
		return
	}

	return
}

func countInv(request inputDocumentListInv) (result []documentInvCount, err error) {

	if err = mssql.DB.Select("COUNT (dh.document_id ) AS document_id_total ,SUM(dh.grand_total ) AS sum_grand_total").
		Table("document_header AS dh, config AS c, company AS c2").
		Where("dh.document_type = 210 AND dh.document_number LIKE ? AND dh.currency IN ( ? ) AND dh.seller_company_id = ? AND dh.seller_company_id = c2.id AND dh.document_status IN ( ? ) AND c.status_code = dh.document_status AND dh.create_date BETWEEN ? and ? AND dh.expire_date BETWEEN ? and ? AND dh.delivery_date BETWEEN ? and ? ", "%"+request.DocumentNumber+"%", request.Currency, request.SellerCompanyID, request.DocumentStatus, request.CreateDate, request.CreateDate2, request.ExpireDate, request.ExpireDate2, request.DeliveryDate, request.DeliveryDate2).
		Find(&result).Error; err != nil {
		return
	}

	return
}

//classify seller or buyer by ID
func getCompanyById(request inputStatus) (result CompanyInfo, err error) {

	if err = mssql.DB.Select("*").
		Table("company").
		Where("id = ?", request.LoginCompId).
		Find(&result).Error; err != nil {
		return
	}
	return
}

//update data to field
func updatedocCancel(request inputStatus, Cancelled int, docTypeCode int64, status []int) (err error) {
	if err = mssql.DB.Table("document_header").
		Where("document_number = ? AND seller_company_id = ? AND buyer_company_id = ? AND document_type = ? AND document_status IN (?)", request.DocumentNumber, request.TargetCompId, request.LoginCompId, docTypeCode, status).
		Update(map[string]interface{}{
			"update_date":       time.Now(),
			"update_by_comp_id": request.LoginCompId,
			"document_status":   Cancelled,
			"reason_cancel":     request.ReasonCancel,
			"record_status":     0,
		}).
		Error; err != nil {
		return
	}
	return
}

//reject update to field
func updatedocReject(request inputStatus, Cancelled, typecode int64, status []int) (err error) { //request is from logic
	if err = mssql.DB.Table("document_header").
		Where("document_number = ? AND seller_company_id = ? AND buyer_company_id = ? AND document_type = ? AND document_status IN (?)", request.DocumentNumber, request.TargetCompId, request.LoginCompId, typecode, status).
		Update(map[string]interface{}{
			"update_date":       time.Now(),
			"update_by_comp_id": request.TargetCompId,
			"document_status":   Cancelled,
			"reason_reject":     request.Note,
		}).
		Error; err != nil {
		return
	}
	return
}
