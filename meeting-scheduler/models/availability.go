package models

import "time"

type UserAvailability struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `json:"event_id"`
	UserID    uint      `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

