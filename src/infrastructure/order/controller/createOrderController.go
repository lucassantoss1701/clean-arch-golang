package controller

import (
	"lucassantoss1701/clean/src/entity/item/models"
	paymentModel "lucassantoss1701/clean/src/entity/payment/models"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	createOrderUseCase *orderUseCase.CreateOrderUseCase
}

func NewCreateOrderController(orderUseCase *orderUseCase.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{
		createOrderUseCase: orderUseCase,
	}
}

func (controller *CreateOrderController) CreateOrder(httpContext *gin.Context) {
	var request CreateOrderRequest
	httpContext.BindJSON(&request)

	input := &orderUseCase.Input{
		Items:       request.Items,
		PaymentType: request.PaymentType,
	}

	controller.createOrderUseCase.Execute(input)

	httpContext.JSON(http.StatusOK, request)
}

type CreateOrderRequest struct {
	Items       []models.Item               `json:"items"`
	PaymentType paymentModel.PAYMENT_METHOD `json:"payment_type"`
}
