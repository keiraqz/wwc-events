package models

import time

// Event represents an event record.
type Event struct {
	Id        int    `json:"id" db:"id"`
	EventName string `json:"eventName" db:"event_name"`
	EventDate time.Time    `json:"eventDate" db:"event_date"`
}

