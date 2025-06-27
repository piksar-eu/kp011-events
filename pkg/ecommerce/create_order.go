package ecommerce

import (
	"events/pkg/ecommerce/public"
	"events/pkg/events"
	"events/pkg/shared"
	"math/rand"

	"github.com/google/uuid"
)

type OrderHandler struct {
	OrderRepository OrderRepository
	EventPublisher  events.EventPublisher
}

func (h *OrderHandler) Create(cmd public.CreateOrderCommand) {
	id := uuid.NewString()
	order := createOrder(id, cmd.Customer, nil)
	for _, i := range cmd.Items {
		order.addItems(OrderItem{
			Ean:      i.Ean,
			Quantity: i.Quantity,
			Price:    calcPrice(i.Ean),
		})
	}

	h.OrderRepository.Save(order)
	// domain events
	h.EventPublisher.Publish(order.PullEvents()...)
	// integration event
	h.EventPublisher.Publish(shared.NewSalesRecordedEvent("ecommerce", id))
}

func calcPrice(ean string) float64 {
	return float64(rand.Intn(99)) / 100
}
