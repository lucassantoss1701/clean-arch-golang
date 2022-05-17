package routes

import (
	"lucassantoss1701/clean/src/infrastructure/order/controller"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, controller *controller.CreateOrderController) {
	router.POST("/orders", controller.CreateOrder)
}
