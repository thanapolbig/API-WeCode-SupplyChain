package onboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{

	}
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) GetCompanyEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingGetPramsEndpoint")

	//ดึงค่าจาก params ชื่อ name
	companyId := c.Params.ByName("id")
	log.Infof("Params company id [%s]", companyId)

	//เรียก logic
	result, err := getCompany(companyId)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//return success
	c.JSON(http.StatusOK, result)
	return
}
