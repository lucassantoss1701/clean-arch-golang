package router

import (
	config "lucassantoss1701/clean/src/infrastructure/config"
	routes "lucassantoss1701/clean/src/infrastructure/order/routes"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"

	"github.com/gin-gonic/gin"
)

type ApplicationContext struct {
	CreateOrderUseCase *orderUseCase.CreateOrderUseCase
}

func SetupHTTP(applicationContext *config.ApplicationContext) *gin.Engine {
	router := gin.Default()

	prepareAllRoutes(router, applicationContext)

	return router
}

func prepareAllRoutes(router *gin.Engine, applicationContext *config.ApplicationContext) {
	routes.SetupOrderRoutes(router, applicationContext.CreateOrderController)
}
