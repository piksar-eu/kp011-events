package ecommerce

import "events/pkg/events"

type OrderItem struct {
	Ean      string  `json:"ean"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Order struct {
	Id        string      `json:"id"`
	Customeer string      `json:"customer"`
	Items     []OrderItem `json:"items"`

	events.DomainEvents
}

func (o *Order) addItems(items ...OrderItem) {
	o.Items = append(o.Items, items...)
}

func createOrder(id, customer string, items []OrderItem) *Order {
	order := &Order{
		Id:        id,
		Customeer: customer,
		Items:     items,
	}

	order.Record(newOrderCreatedEvent(id))

	return order
}

type OrderRepository interface {
	Save(*Order)
}

const orderCreatedEvent = "ecommerce.OrderCreatedEvent"

type orderCreatedPayload struct {
	OrderId string `json:"orderId"`
}

func newOrderCreatedEvent(oId string) events.Event {
	return events.NewEvent(orderCreatedEvent, orderCreatedPayload{
		OrderId: oId,
	})
}
