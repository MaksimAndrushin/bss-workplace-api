package business

import "fmt"

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Locked
	Processed
)

type WorkplaceEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Workplace
}

func (e *WorkplaceEvent) String() string {
	return fmt.Sprintf("Event: ID - %d, Type - %d, Status - %d, Entity - %v",
		e.ID,
		e.Type,
		e.Status,
		e.Entity)
}
