package documents

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getDocumentDetail(request documentDetailReq) (result documentDetailRes, err error) {
	var detail []documentDetail
	detail, err = getDocDetail(request)
	if err != nil {
		return
	}
	log.Infof("document_detail : %+v", detail)

	var header documentHeader
	header, err = getDocHeader(request, nil)
	if err != nil {
		return
	}
	log.Infof("document_detail : %+v", header)

	result = documentDetailRes{
		header,
		detail,
	}
	return
}

//createDocument
func createDocument(request draftDocument) (result messageResponse, err error) {

	// set  status and date
	request.Header.DocumentStatus = Draft
	request.Header.RecordStatus = true
	// date
	request.Header.CreateDate = time.Now()
	current := time.Now().Format("2006-01-02")
	currentStr := strings.Join(strings.Split(current, "-"), "")

	//check documentNumber
	var checkDocHeader draftDocumentHeader
	switch request.Header.DocumentType {
	case 110:
		request.Header.DocumentNumber = "PO" + currentStr + "_"
		checkDocHeader, err = checkDocumentHeader(request.Header)
		if err != nil && err.Error() != "record not found"{
			return
		}

	case 120:
		request.Header.DocumentNumber = "GR" + currentStr + "_"
		checkDocHeader, err = checkDocumentHeader(request.Header)
		if err != nil && err.Error() != "record not found"{
			return
		}

	case 210:
		request.Header.DocumentNumber = "INV" + currentStr + "_"
		checkDocHeader, err = checkDocumentHeader(request.Header)
		if err != nil && err.Error() != "record not found"{
			return
		}
	}

	var headerRes messageResponse
	// record not found
	if err != nil {

		//create createDocumentHeader
		if err.Error() == "record not found" {
			request.Header.DocumentNumber = request.Header.DocumentNumber + "0000"
			headerRes, err = createDocumentHeader(request.Header)
			if err != nil {
				return
			}
		} else {
			return
		}
	} else {
		// getDocument Number + 1
		documentNumberInt, _ := strconv.Atoi(checkDocHeader.DocumentNumber[11:len(checkDocHeader.DocumentNumber)])
		documentNumberStr := PadLeft(strconv.Itoa(documentNumberInt+1), "0", 4)

		//create createDocumentHeader
		request.Header.DocumentNumber = request.Header.DocumentNumber + documentNumberStr
		headerRes, err = createDocumentHeader(request.Header)
		if err != nil {
			return
		}
	}
	// struct
	var documentReq = documentDetailReq{
		DocumentNumber:  request.Header.DocumentNumber,
		SellerCompanyId: request.Header.SellerCompanyId,
		BuyerCompanyId:  request.Header.BuyerCompanyId,
		DocumentType:    request.Header.DocumentType,
	}
	//get documentId
	var docHeader documentHeader
	docHeader, err = getDocHeader(documentReq, nil)
	if err != nil {
		return
	}
	// set documentId in document_detail
	for index := 0 ; index < len(request.Detail) ; index ++ {
		request.Detail[index].DocumentHeaderId = docHeader.DocumentId
	}
	log.Infof("document_detail : %+v", docHeader)
	//create document detail
	var detailRes messageResponse
	detailRes, err = createDocumentDetail(request.Detail)
	if err != nil {
		return
	}

	if detailRes.Status == 201 && headerRes.Status == 201 {
		result.Status = http.StatusCreated
		result.MessageCode = "0000"
		result.MessageDescription = "create " + docHeader.DocumentNumber[0:2] + "[" + docHeader.DocumentNumber + "]"
	}
	return
}

//func generateDocumentNumber(request draftDocumentHeader) (result documentDetailRes, err error){
//	return
//}
func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[1 : lenght+1]
		}
	}
}

func getPODocumentList(request documentListReq) (result documentListRes, err error) {
	var list []documentList
	list, err = getPODocList(request)
	if err != nil {
		return
	}
	log.Infof("documentList : %+v", list)

	var count PoCount
	count, err = getPODocListCount(request)
	if err != nil {
		return
	}
	log.Infof("documentListCount : %+v", list)

	result = documentListRes{
		count,
		list,
		// header,
	}
	return
}

func getDocumentListInv(request inputDocumentListInv) (result documentInvRes, err error) {
	var detailInv []documentListInv
	detailInv, err = documentListInvGet(request)
	if err != nil {
		return
	}
	log.Infof("documentDetailInv : %+v", detailInv)

	var count []documentInvCount
	count, err = countInv(request)
	if err != nil {
		return
	}
	log.Infof("documentDetailInv : %+v", count)

	result = documentInvRes{
		count,
		detailInv,
	}

	return
}

//cancel document
func GetCancel(request inputStatus, typecode int64) (mes MessageResponse, err error) {
	var comp CompanyInfo
	comp, err = getCompanyById(request)
	if err != nil {
		return
	} //classify seller or buyer

	if typecode == PurchaseOrder && typecode == GoodsReceipt { //buyer's allow to see only PO and GR
		if comp.CompanyType != Seller {
			return
		}
	}

	if typecode == Invoice { //INV only seller's allow to see
		if comp.CompanyType != Buyer {
			return
		}
	}

	statusDoc := []int{Draft, ForModify, ForModifyAfterSubmit}
	documentDetailReq := documentDetailReq{
		DocumentNumber:  request.DocumentNumber,
		SellerCompanyId: request.TargetCompId,
		BuyerCompanyId:  request.LoginCompId,
		DocumentType:    typecode,
	}

	_, err = getDocHeader(documentDetailReq, statusDoc) //type of document
	if err != nil {
		return
	}

	//update data to field
	err = updatedocCancel(request, Cancelled, typecode, statusDoc)
	if err != nil {
		return
	}

	mes = MessageResponse{
		Status:             http.StatusOK,
		MessageCode:        "0000",
		MessageDescription: "update data success",
	}
	return

}
