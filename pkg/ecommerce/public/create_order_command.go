package public

type OrderItem struct {
	Ean      string
	Quantity int
}

type CreateOrderCommand struct {
	Customer string
	Items    []OrderItem
}
