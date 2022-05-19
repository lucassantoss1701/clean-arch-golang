package gateway

import (
	"database/sql"
	"lucassantoss1701/clean/src/entity/order/models"
)

type OrderDbGateway struct {
	conn *sql.DB
}

func NewOrderDbGateway(db *sql.DB) *OrderDbGateway {
	return &OrderDbGateway{
		conn: db,
	}
}

func (orderDbGateway *OrderDbGateway) Create(order *models.Order) error {
	statement, err := orderDbGateway.conn.Prepare("INSERT INTO `order` (payment_method) VALUES(?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(order.PaymentMethod)
	if err != nil {
		return err
	}

	return nil
}
