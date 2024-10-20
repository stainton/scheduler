package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stainton/scheduler/internal/database"
	"github.com/stainton/scheduler/internal/handlers"
	"github.com/stainton/scheduler/internal/scheduler"
)

func main() {
	database.InitDB()

	go scheduler.ScheduleTasks()
	go scheduler.MonitorTasks()

	r := gin.Default()
	r.POST("/tasks", handlers.TaskCreationHandler)
	r.POST("/tasks/result", handlers.TaskResultHandler)
	r.Run(":8080")
}
