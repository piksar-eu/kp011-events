package shared

import "events/pkg/events"

const SalesRecordedEvent = "shared.SalesRecordedEvent"

type SalesRecordedPayload struct {
	Source string `json:"source"`
	Id     string `json:"id"`
}

func NewSalesRecordedEvent(source, id string) events.Event {
	return events.NewEvent(SalesRecordedEvent, SalesRecordedPayload{
		Source: source,
		Id:     id,
	})
}
