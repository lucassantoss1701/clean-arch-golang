package routes

import (
	"lucassantoss1701/clean/src/infrastructure/order/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutesOrder(r *gin.Engine, controller *controller.CreateOrderController) {
	r.POST("/orders", controller.CreateOrder)
}
