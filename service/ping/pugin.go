package ping

import (
	mssql "api-wecode-supplychain/database/mssql"
)

func logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = mssql.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}
