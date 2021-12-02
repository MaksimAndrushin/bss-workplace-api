package model

func CreateEventFromWorkplace(eventType EventType, eventStatus EventStatus, workplace Workplace) *WorkplaceEvent {
	workplaceEntity := WorkplaceEvent{
		Type:        eventType,
		Status:      eventStatus,
		Entity:      &workplace,
	}

	return &workplaceEntity
}
