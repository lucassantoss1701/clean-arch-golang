package gateway

import (
	"database/sql"
	orderModel "lucassantoss1701/clean/src/entity/order/model"
)

type OrderDbGateway struct {
	conn *sql.DB
}

func NewOrderDbGateway(db *sql.DB) *OrderDbGateway {
	return &OrderDbGateway{
		conn: db,
	}
}

func (orderDbGateway *OrderDbGateway) Create(order *orderModel.Order) error {

	statement, err := orderDbGateway.conn.Prepare("INSERT INTO `order` (payment_method) VALUES(?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	savedOrder, err := statement.Exec(order.PaymentMethod)

	if err != nil {
		return err
	}

	orderId, err := savedOrder.LastInsertId()
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		statement, err = orderDbGateway.conn.Prepare("INSERT INTO `order_item` (order_id, item_id, quantity, price) VALUES(?,?,?,?)")

		if err != nil {
			return err
		}

		_, err = statement.Exec(orderId, item.Id, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	return nil
}
