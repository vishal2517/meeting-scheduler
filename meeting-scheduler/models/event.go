package models

import "time"

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Duration  int       `json:"duration"` // in minutes
	CreatedAt time.Time `json:"created_at"`
}

type EventSlot struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `json:"event_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

