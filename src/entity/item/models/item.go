package models

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

func NewItem(id int, name string, quantity int, price int) *Item {
	return &Item{
		Id:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
}

func (item *Item) Total() int {
	return item.Price * item.Quantity
}
