package onboard

import (
	mssql "api-wecode-supplychain/database/mssql"
)

func getCompanyById(companyId string) (result companyInfo, err error) {

	//if err = mssql.DB.Select("c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province").
	//	Table("company AS c").
	//	Joins("LEFT JOIN config AS cf ON c.company_type = cf.status_code").
	//	Where("id=?", companyId).
	//	Find(&result).Error; err != nil {
	//	return
	//}

	if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
		return
	}

	return

}
