package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stainton/scheduler/internal/models"
)

type TaskResult struct {
	TaskID uint   `json:"task_id"`
	Status string `json:"status"`
}

func TaskResultHandler(c *gin.Context) {
	var result TaskResult

	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	tasks, err := models.GetTaskByID(result.TaskID)
	if err != nil || len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	task := tasks[0]

	task.Status = result.Status
	task.UpdatedAt = time.Now()
	err = models.UpdateTaskStatus(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task status updated"})
}
