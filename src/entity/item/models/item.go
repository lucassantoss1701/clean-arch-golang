package models

type Item struct {
	id       int
	name     string
	quantity int
	price    int
}

func NewItem(id int, name string, quantity int, price int) *Item {
	return &Item{
		id:       id,
		name:     name,
		quantity: quantity,
		price:    price,
	}
}

func (i *Item) Total() int {
	return i.price * i.quantity
}
