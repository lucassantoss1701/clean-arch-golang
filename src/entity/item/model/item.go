package model

type Item struct {
	Id       int
	Name     string
	Quantity int
	Price    int
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
