package controller

import (
	"lucassantoss1701/clean/src/entity/item/models"
	paymentModel "lucassantoss1701/clean/src/entity/payment/models"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	ItemsDTO    []*ItemDTO                `json:"items"`
	PaymentType paymentModel.PAYMENT_TYPE `json:"payment_type"`
}

type CreateOrderController struct {
	createOrderUseCase *orderUseCase.CreateOrderUseCase
}

func toEntity(itemsDTO []*ItemDTO) []models.Item {
	var items []models.Item
	for _, itemDTO := range itemsDTO {
		item := models.NewItem(itemDTO.Id, itemDTO.Name, itemDTO.Quantity, itemDTO.Price)
		items = append(items, *item)
	}

	return items
}

func NewCreateOrderController(orderUseCase *orderUseCase.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{
		createOrderUseCase: orderUseCase,
	}
}

func (createOrderController *CreateOrderController) CreateOrder(c *gin.Context) {
	var request CreateOrderRequest
	c.BindJSON(&request)

	input := &orderUseCase.Input{
		Items:       toEntity(request.ItemsDTO),
		PaymentType: request.PaymentType,
	}

	createOrderController.createOrderUseCase.Execute(input)

	c.JSON(http.StatusOK, request)
}
