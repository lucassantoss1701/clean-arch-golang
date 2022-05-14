package router

import (
	"lucassantoss1701/clean/src/infrastructure/config"
	"lucassantoss1701/clean/src/infrastructure/order/routes"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"

	"github.com/gin-gonic/gin"
)

type ApplicationContext struct {
	CreateOrderUseCase *orderUseCase.CreateOrderUseCase
}

func SetupHTTP(applicationContext *config.ApplicationContext) *gin.Engine {
	r := gin.Default()
	prepareAllRoutes(r, applicationContext)
	return r
}

func prepareAllRoutes(r *gin.Engine, applicationContext *config.ApplicationContext) {
	routes.SetupRoutesOrder(r, applicationContext.CreateOrderController)
}
