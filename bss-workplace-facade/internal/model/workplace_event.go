package model

import "time"

type EventType uint32

type EventStatus uint32

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
	WorkplaceId uint64
	Type        EventType
	Status      EventStatus
	Updated     time.Time
	Entity      *Workplace
}
