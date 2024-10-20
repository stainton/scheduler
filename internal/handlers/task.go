package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stainton/scheduler/internal/models"
)

func TaskCreationHandler(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	task.Status = "pending"
	task.CreatedAt = time.Now()
	task.ScheduledAt = time.Now().Add(10 * time.Minute)

	if err := models.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task create successfully",
		"task":    task,
	})
}
