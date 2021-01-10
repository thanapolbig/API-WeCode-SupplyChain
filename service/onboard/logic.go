package onboard

func getCompany(companyId string) (result companyInfo, err error){

	result, err = getCompanyById(companyId)
	if err != nil {
		return
	}

	return
}

