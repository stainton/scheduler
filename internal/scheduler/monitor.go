package scheduler

import (
	"log"
	"time"

	"github.com/stainton/scheduler/internal/models"
)

func MonitorTasks() {
	for {
		tasks, err := models.GetRunningTasks()
		if err != nil {
			log.Println("error fetching tasks:", err)
			continue
		}

		for _, task := range tasks {
			if time.Since(task.UpdatedAt) > 30*time.Minute {
				task.Status = "timeout"
				err := models.UpdateTaskStatus(task)
				if err != nil {
					log.Println("error updating task status:", err)
				}
				log.Printf("task %d failed due to timeout", task.ID)
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
