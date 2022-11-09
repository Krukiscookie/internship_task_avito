package handler

import (
	_ "github.com/Krukiscookie/intern_task/docs"
	"github.com/Krukiscookie/intern_task/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	balance := router.Group("/user-money")
	{
		balance.GET("/:id", h.GetBalance)
		balance.POST("/transfer", h.TransferMoney)
		balance.POST("/addmoney", h.AddMoney)
		balance.POST("/services/reserve", h.ServiceReserve)
		balance.POST("/services/approve", h.ServiceApprove)
		balance.POST("/services/refusal", h.ServiceRefusal)
	}

	reports := router.Group("/reports")
	{
		reports.POST("/transaction", h.TransactionInfo)
		reports.POST("/services-report", h.ServiceReport)
		reports.GET("/:path", h.ReadFile)
	}

	return router
}
