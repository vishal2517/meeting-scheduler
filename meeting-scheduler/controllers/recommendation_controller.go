package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"meeting-scheduler/db"
	"meeting-scheduler/models"
	"net/http"
	//"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRecommendedSlots godoc
// @Summary Get recommended time slots
// @Description Returns the best meeting slot for all users
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Router /events/{id}/recommendations [get]
func GetRecommendedSlots(c *gin.Context) {
	eventID := c.Param("id")

	ctx := context.Background()
	cacheKey := "recommended_slots:" + eventID

	// Check if result exists in Redis
	cachedData, err := db.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("✅ Returning cached result from Redis")
		var response map[string]interface{}
		json.Unmarshal([]byte(cachedData), &response)
		c.JSON(http.StatusOK, response)
		return
	}

	// If not cached, fetch data from DB
	var event models.Event
	if err := db.DB.Where("id = ?", eventID).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var slots []models.EventSlot
	db.DB.Where("event_id = ?", eventID).Find(&slots)

	var availabilities []models.UserAvailability
	db.DB.Where("event_id = ?", eventID).Find(&availabilities)

	// Compute best time slot
	slotCount := make(map[uint]int)
	for _, slot := range slots {
		for _, avail := range availabilities {
			if avail.StartTime.Before(slot.EndTime) && avail.EndTime.After(slot.StartTime) {
				slotCount[slot.ID]++
			}
		}
	}

	// Find the slot with maximum users
	var bestSlot models.EventSlot
	maxUsers := 0
	for _, slot := range slots {
		if slotCount[slot.ID] > maxUsers {
			bestSlot = slot
			maxUsers = slotCount[slot.ID]
		}
	}

	// Prepare response
	response := gin.H{
		"event_id": eventID,
		"best_slot": gin.H{
			"id":        bestSlot.ID,
			"start":     bestSlot.StartTime,
			"end":       bestSlot.EndTime,
		},
		"max_users": maxUsers,
	}

	// Cache the result in Redis for 10 minutes
	cacheData, _ := json.Marshal(response)
	db.RedisClient.Set(ctx, cacheKey, cacheData, 10*time.Minute)

	fmt.Println("✅ Cached meeting slot recommendations in Redis")

	c.JSON(http.StatusOK, response)
}
