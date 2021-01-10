package documents

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

//รับ INPUT แปลงค่า
//document
func (ep *Endpoint) GetDocumentDetailEndpoint(c *gin.Context) { //POST documents/getDocumentDetail
	defer c.Request.Body.Close()
	log.Infof("Get DocumentDetail Endpoint")


	//model รับ input จาก body
	var request documentDetailReq
	if err := c.ShouldBindBodyWith(&request, binding.JSON);
	err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("request : %+v", request)

	// รับ param
	docType := strings.ToLower(c.Params.ByName("type"))
	log.Infof("Params type : %s", docType)

	switch docType {
	case "po":
		request.DocumentType = PurchaseOrder
		//เรียก logic
		result, err := getDocumentDetail(request)
		if err != nil {
			//not found
			if( err.Error() == "record not found"){
				message := MessageResponse{ //default response
					Status:             http.StatusNoContent,
					MessageCode:        "00000",
					MessageDescription: "document NOT FOUND",
				}
				c.JSON(message.Status, message)
				return
			}
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	case "inv":
		request.DocumentType = Invoice
		//เรียก logic
		result, err := getDocumentDetail(request)
		if err != nil {
			//not found
			if( err.Error() == "record not found"){
				message := MessageResponse{ //default response
					Status:             http.StatusNoContent,
					MessageCode:        "00000",
					MessageDescription: "document NOT FOUND",
				}
				c.JSON(message.Status, message)
				return
			}
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	case "gr":
		request.DocumentType = GoodsReceipt
		//เรียก logic
		result, err := getDocumentDetail(request)
		if err != nil {
			//not found
			if( err.Error() == "record not found"){
				message := MessageResponse{ //default response
					Status:             http.StatusNoContent,
					MessageCode:        "00000",
					MessageDescription: "document NOT FOUND",
				}
				c.JSON(message.Status, message)
				return
			}
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	default:
		message := MessageResponse{ //default response
			Status:             http.StatusNotFound,
			MessageCode:        "00000",
			MessageDescription: "URL NOT FOUND",
		}
		c.JSON(message.Status, message)
		return
	}
}

func (ep *Endpoint) DraftDocumentDetailEndpoint(c *gin.Context) { //POST documents/:type/draft
	defer c.Request.Body.Close()
	log.Infof("Get DocumentDetail Endpoint")

	//model รับ input จาก body
	var request draftDocument
	if err := c.ShouldBindBodyWith(&request, binding.JSON);
		err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("request : %+v", request)

	// รับ param
	docType := strings.ToLower(c.Params.ByName("type"))
	log.Infof("Params type : %s", docType)

	switch docType {
	case "po":
		// set type
		request.Header.DocumentType = PurchaseOrder

		//เรียก logic
		result, err := createDocument(request)
		if err != nil {
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	case "inv":
		// set type
		request.Header.DocumentType = Invoice
		//เรียก logic
		result, err := createDocument(request)
		if err != nil {
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	case "gr":
		// set type
		request.Header.DocumentType = GoodsReceipt
		//เรียก logic
		result, err := createDocument(request)
		if err != nil {
			//return err
			message := MessageResponse{ //default response
				Status:             http.StatusBadRequest,
				MessageCode:        "00000",
				MessageDescription: err.Error(),
			}
			c.JSON(message.Status, message)
			return
		}
		//return success
		c.JSON(http.StatusOK, result)
		return
	default:
		message := MessageResponse{ //default response
			Status:             http.StatusNotFound,
			MessageCode:        "00000",
			MessageDescription: "URL NOT FOUND",
		}
		c.JSON(message.Status, message)
		return
	}
}

func (ep *Endpoint) GettPurchesOrderDocListEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Get DocumentList Endpoint")

	var request documentListReq
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("request : %+v", request)

	result, err := getPODocumentList(request)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) DocumentListInv(c *gin.Context) { //POST /documents/getDocumentList
	defer c.Request.Body.Close()
	log.Infof("DocumentListINV")

	//ดึงค่าจาก body
	var request inputDocumentListInv //model รับ input จาก body

	//เรียก logic
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("request : %+v", request)

	//เรียก logic
	result, err := getDocumentListInv(request)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//return success
	c.JSON(http.StatusOK, result)
	return
}

//cancle document
func (ep *Endpoint) CancelDocumentType(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Cancel document")

	document := c.Params.ByName("type") //input type in parameter
	log.Infof("Type : %s", document)    //input type PO GR or INV

	var request inputStatus            //model รับ input จาก body
	var typecode int64                 //typecode
	switch strings.ToUpper(document) { //case by case of PO GR INV from parameter
	case "PO":
		typecode = PurchaseOrder
	case "GR":
		typecode = GoodsReceipt
	case "INV":
		typecode = Invoice
	default:
		message := MessageResponse{ //default response
			Status:             http.StatusNotFound,
			MessageCode:        "00000",
			MessageDescription: "url not found",
		} //return 404 not found status
		c.JSON(message.Status, message)
		return
	}

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} //error with 400

	log.Infof("request login company ID : %+v", request.LoginCompId)
	log.Infof("request target company ID : %+v", request.TargetCompId)
	log.Infof("request document number : %+v", request.DocumentNumber)
	log.Infof("request reason : %+v", request.ReasonCancel)

	result, err := GetCancel(request, typecode) //เรียก logic มาใช้
	if err != nil {
		mes := MessageResponse{
			Status:             http.StatusBadRequest,
			MessageCode:        "00000",
			MessageDescription: "fail to update data",
		}
		c.JSON(mes.Status, mes)
		return
	}

	c.JSON(http.StatusOK, result)
	return

}
