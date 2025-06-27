package main

import (
	"events/internal"
	"events/pkg/ecommerce"
	ecommerce_public "events/pkg/ecommerce/public"
	"events/pkg/shared"
	"events/pkg/stock"
)

func main() {

	eventPublisher := internal.NewSyncEventPublisher()

	eventPublisher.Register(shared.SalesRecordedEvent, stock.ReserveStockOnSalesRecordedListener)

	ecommerceOrderHandler := ecommerce.OrderHandler{
		OrderRepository: &internal.EcommerceOrderRepository{},
		EventPublisher:  eventPublisher,
	}

	cmd := ecommerce_public.CreateOrderCommand{
		Customer: "John Doe",
		Items: []ecommerce_public.OrderItem{
			{Ean: "2065126418793", Quantity: 6},
			{Ean: "5589848367490", Quantity: 3},
		},
	}

	ecommerceOrderHandler.Create(cmd)
}
