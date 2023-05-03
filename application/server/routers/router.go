package routers

import (
	v1 "application/api/v1"
	v2 "application/api/v2"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)
		apiV1.POST("/queryAccountList", v1.QueryAccountList)
		apiV1.POST("/createRealEstate", v1.CreateRealEstate)
		apiV1.POST("/queryRealEstateList", v1.QueryRealEstateList)
		apiV1.POST("/createSelling", v1.CreateSelling)
		apiV1.POST("/createSellingByBuy", v1.CreateSellingByBuy)
		apiV1.POST("/querySellingList", v1.QuerySellingList)
		apiV1.POST("/querySellingListByBuyer", v1.QuerySellingListByBuyer)
		apiV1.POST("/updateSelling", v1.UpdateSelling)
		apiV1.POST("/createDonating", v1.CreateDonating)
		apiV1.POST("/queryDonatingList", v1.QueryDonatingList)
		apiV1.POST("/queryDonatingListByGrantee", v1.QueryDonatingListByGrantee)
		apiV1.POST("/updateDonating", v1.UpdateDonating)
	}
	apiV2 := r.Group("/api/v2")
	{
		apiV2.GET("/hello", v2.Hello)
		apiV2.POST("/queryAccountV2List", v2.QueryAccountV2List)
		apiV2.POST("/createPrescription", v2.CreatePrescription)
		apiV2.POST("/queryPrescription", v2.QueryPrescriptionList)
	}
	return r
}
