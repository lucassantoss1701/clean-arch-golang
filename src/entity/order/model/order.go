package model

import (
	itemModel "lucassantoss1701/clean/src/entity/item/model"
	paymentModel "lucassantoss1701/clean/src/entity/payment/model"

	time "time"
)

type Order struct {
	Cupom         int
	Items         []itemModel.Item
	CreatedAt     time.Time
	PaymentMethod paymentModel.PAYMENT_METHOD
	Code          int
}

func NewOrder(items []itemModel.Item, paymentMethod paymentModel.PAYMENT_METHOD) *Order {
	return &Order{
		Items:         items,
		CreatedAt:     time.Now(),
		PaymentMethod: paymentMethod,
	}
}

func (order *Order) Total() int {
	total := 0
	for _, item := range order.Items {
		total = +item.Total()
	}
	return total
}
