package controller

import (
	"fmt"
	itemModel "lucassantoss1701/clean/src/entity/item/model"
	paymentModel "lucassantoss1701/clean/src/entity/payment/model"
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

	fmt.Println("CreateOrderController.CreateOrderRequest", request)

	controller.createOrderUseCase.Execute(request.toInput())

	httpContext.JSON(http.StatusOK, request)
}

type CreateOrderRequest struct {
	Items         []CreateOrderItemRequest    `json:"items"`
	PaymentMethod paymentModel.PAYMENT_METHOD `json:"payment_method"`
}

func (request *CreateOrderRequest) toInput() *orderUseCase.Input {
	var items []itemModel.Item

	for _, item := range request.Items {
		items = append(items, *itemModel.NewItem(item.Id, item.Name, item.Quantity, item.Price))
	}

	return &orderUseCase.Input{
		Items:         items,
		PaymentMethod: request.PaymentMethod,
	}
}

type CreateOrderItemRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
