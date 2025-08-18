package stock

import (
	"events/pkg/events"
	"events/pkg/shared"
	"log"
)

func ReserveStockOnSalesRecordedListener(e events.Event) {
	if e.Type != shared.SalesRecordedEvent {
		log.Println("invalid event type")
		return
	}

	payload := e.Payload.(shared.SalesRecordedPayload)

	log.Println(payload)
}
