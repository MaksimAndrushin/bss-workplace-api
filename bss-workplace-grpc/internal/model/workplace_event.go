package model

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
	ID          uint64
	Type        EventType
	Status      EventStatus
	Entity      *Workplace
}
