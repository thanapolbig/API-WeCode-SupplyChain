package onboard

type companyInfo struct {
	Id              int    `json:"id"`
	CompanyName     string `json:"company_name"`
	ShortName       string `json:"short_name"`
	CompanyTypeCode int    `json:"company_type_code"`
	CompanyTypeName string `json:"company_type_name"`
	Branch          string `json:"branch"`
	PhoneNumber     string `json:"phone_number"`
	District        string `json:"district"`
	SubDistrict     string `json:"sub_district"`
	ZipCode         string `json:"zip_code"`
	Province        string `json:"province"`
}
