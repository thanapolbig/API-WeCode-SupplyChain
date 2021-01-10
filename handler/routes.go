package handler

import (
	"net/http"

	"api-wecode-supplychain/service/dashboard"
	"api-wecode-supplychain/service/documents"
	"api-wecode-supplychain/service/onboard"
	"api-wecode-supplychain/service/ping"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	ping := ping.NewEndpoint()
	dashboard := dashboard.NewEndpoint()
	documents := documents.NewEndpoint()
	onboard := onboard.NewEndpoint()

	r.transaction = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/:type/ping",
			Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/:type/ping/:name",
			Endpoint:    ping.PingGetParamsEndpoint,
		},
		{
			Name:        "Ping Pong : POST Prams+Body",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodPost,
			Pattern:     "/:type/ping/:name",
			Endpoint:    ping.PingPostParamsAndBodyEndpoint,
		},
	}

	txDashboard := []route{
		{
			Name:        "Buyer : POST Dashboard",
			Description: "Buyer : Data",
			Method:      http.MethodPost,
			Pattern:     "/:type/sumDocumentHeader",
			Endpoint:    dashboard.GetDashboard,
		},
	}

	txDocument := []route{
		{
			Name:        "GET DocumentDetail Endpoint",
			Description: "GET DocumentDetail Endpoint",
			Method:      http.MethodPost,
			Pattern:     "/:type/getDocumentDetail",
			Endpoint:    documents.GetDocumentDetailEndpoint,
		},
		{
			Name:        "GET DocumentListPO Endpoint",
			Description: "GET DocumentListPO Endpoint",
			Method:      http.MethodPost,
			Pattern:     "/:type/getDocumentList",
			Endpoint:    documents.GettPurchesOrderDocListEndpoint,
		},
		{
			Name:        "Post DocumentListINV",
			Description: "Post DocumentListINV",
			Method:      http.MethodPost,
			Pattern:     "/:type/getDocumentListInv",
			Endpoint:    documents.DocumentListInv,
		},
		{
			Name:        "Post CancelDocument",
			Description: "Post CancelDocument",
			Method:      http.MethodPost,
			Pattern:     "/:type/cancel",
			Endpoint:    documents.CancelDocumentType, //cancle document
		},
		{
			Name:        "POST DocumentDetail Endpoint",
			Description: "POST create DocumentDetail Endpoint",
			Method:      http.MethodPost,
			Pattern:     "/:type/draft",
			Endpoint:    documents.DraftDocumentDetailEndpoint,
		},

	}

	txOnboard := []route{
		{
			Name:        "Get CompanyInfo Endpoint",
			Description: "Get CompanyInfo Info By Id",
			Method:      http.MethodGet,
			Pattern:     "/getCompanyInfo/:id",
			Endpoint:    onboard.GetCompanyEndpoint,
		},
	}

	ro := gin.New()

	//ro.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/documents")
	for _, e := range txDocument {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/dashboard")
	for _, e := range txDashboard {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/onboard")
	for _, e := range txOnboard {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
