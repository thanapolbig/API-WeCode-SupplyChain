package dashboard

import (
	"time"
)

func checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "Pong"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}

func GetDashboard(request inputDashboard) (result []outputDashboard, company []CompanyInfo, err error) {
	company, err = getCompanyById(request)
	if err != nil {
		return
	}

	for i := 0; i <= len(company); i++ {
		if company[i].CompanyType == Seller {
			result, err = getSellserDashboard(request)
			if err != nil {
				return
			}
			return
		}
		if company[i].CompanyType == Buyer {
			result, err = buyerDashboard(request)
			if err != nil {
				return
			}
			return
		}
	}
	return
}
