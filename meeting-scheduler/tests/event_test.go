package tests

import (
	//"bytes"
	"encoding/json"
	"meeting-scheduler/controllers"
	"meeting-scheduler/db"
	"meeting-scheduler/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup Test Database (Mock)
func setupTestDB() {
	db.ConnectDatabase()
	db.DB.Exec("DELETE FROM events")
	db.DB.Exec("DELETE FROM event_slots")
	db.DB.Exec("DELETE FROM user_availabilities")
}

// Helper function to create an event
func createEvent(title string, duration int) models.Event {
	event := models.Event{Title: title, Duration: duration}
	db.DB.Create(&event)
	return event
}

// Helper function to create an event slot
func createEventSlot(eventID uint, start, end time.Time) models.EventSlot {
	slot := models.EventSlot{EventID: eventID, StartTime: start, EndTime: end}
	db.DB.Create(&slot)
	return slot
}

// Helper function to add user availability
func createUserAvailability(eventID, userID uint, start, end time.Time) models.UserAvailability {
	availability := models.UserAvailability{EventID: eventID, UserID: userID, StartTime: start, EndTime: end}
	db.DB.Create(&availability)
	return availability
}

// Test case for recommending time slots
func TestGetRecommendedSlots(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)

	// Create test event
	event := createEvent("Team Sync", 60)

	// Define time slots for the event
	startTime1, _ := time.Parse(time.RFC3339, "2025-01-12T14:00:00Z")
	endTime1, _ := time.Parse(time.RFC3339, "2025-01-12T16:00:00Z")
	slot1 := createEventSlot(event.ID, startTime1, endTime1)

	//startTime2, _ := time.Parse(time.RFC3339, "2025-01-14T18:00:00Z")
	//endTime2, _ := time.Parse(time.RFC3339, "2025-01-14T21:00:00Z")
	//slot2 := createEventSlot(event.ID, startTime2, endTime2)

	// Add user availabilities
	//user1 := createUserAvailability(event.ID, 1, startTime1, endTime1) // User 1 available in slot 1
	//user2 := createUserAvailability(event.ID, 2, startTime2, endTime2) // User 2 available in slot 2
	//user3 := createUserAvailability(event.ID, 3, startTime1, endTime1) // User 3 available in slot 1

	// Setup router
	router := gin.Default()
	router.GET("/events/:id/recommendations", controllers.GetRecommendedSlots)

	// Test API call
	req, _ := http.NewRequest("GET", "/events/1/recommendations", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Parse response
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotNil(t, response["best_slot"])

	// Check if the recommended slot is the one with the most users
	bestSlot := response["best_slot"].(map[string]interface{})
	assert.Equal(t, float64(slot1.ID), bestSlot["id"])
	assert.Equal(t, float64(2), response["max_users"]) // Slot 1 has 2 users available
}
