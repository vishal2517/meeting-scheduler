package controllers

import (
	"meeting-scheduler/db"
	"meeting-scheduler/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create a new event
// @Description Creates a new meeting event
// @Tags events
// @Accept json
// @Produce json
// @Param event body models.Event true "Event Data"
// @Success 201 {object} models.Event
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&event)
	c.JSON(http.StatusCreated, event)
}

func GetEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := db.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := db.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	db.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}


// Add user availability for an event
func AddUserAvailability(c *gin.Context) {
	var availability models.UserAvailability

	if err := c.ShouldBindJSON(&availability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&availability)
	c.JSON(http.StatusCreated, availability)
}

// GetRecommendedSlots godoc
// @Summary Get recommended time slots
// @Description Returns the best meeting slot for all users
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Router /events/{id}/recommendations [get]
// func GetRecommendedSlots(c *gin.Context) {
// 	eventID := c.Param("id")

// 	// Fetch event slots
// 	var eventSlots []models.EventSlot
// 	db.DB.Where("event_id = ?", eventID).Find(&eventSlots)

// 	// Fetch user availabilities
// 	var availabilities []models.UserAvailability
// 	db.DB.Where("event_id = ?", eventID).Find(&availabilities)

// 	if len(eventSlots) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "No available slots for this event"})
// 		return
// 	}

// 	// Map to count user availability per slot
// 	slotAvailability := make(map[uint]int)
// 	slotUsers := make(map[uint][]uint) // Stores unavailable users per slot

// 	// Check for perfect matches
// 	for _, slot := range eventSlots {
// 		matchingUsers := 0
// 		var unavailableUsers []uint

// 		for _, avail := range availabilities {
// 			if avail.StartTime.Before(slot.EndTime) && avail.EndTime.After(slot.StartTime) {
// 				matchingUsers++
// 			} else {
// 				unavailableUsers = append(unavailableUsers, avail.UserID)
// 			}
// 		}

// 		slotAvailability[slot.ID] = matchingUsers
// 		slotUsers[slot.ID] = unavailableUsers
// 	}

// 	// Find the best slot with maximum availability
// 	var bestSlot models.EventSlot
// 	maxUsers := 0
// 	var unavailableUsers []uint

// 	for _, slot := range eventSlots {
// 		if slotAvailability[slot.ID] > maxUsers {
// 			maxUsers = slotAvailability[slot.ID]
// 			bestSlot = slot
// 			unavailableUsers = slotUsers[slot.ID]
// 		}
// 	}

// 	// Return recommended slot
// 	c.JSON(http.StatusOK, gin.H{
// 		"best_slot":        bestSlot,
// 		"max_users":        maxUsers,
// 		"unavailable_users": unavailableUsers,
// 	})
// }
