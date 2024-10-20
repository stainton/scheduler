package scheduler

import (
	"log"
	"time"

	"github.com/stainton/scheduler/internal/models"
)

func dispatchTaskToWorker(task *models.Task) error {
	log.Printf("Dispatch task %d to worker: %s", task.ID, task.Command)
	return nil
}

func ScheduleTasks() {
	for {
		tasks, err := models.GetPendingTasks()
		if err != nil {
			log.Println("Error fetching pending tasks: ", err)
			continue
		}

		for _, task := range tasks {
			if time.Now().After(task.ScheduledAt) {
				err := dispatchTaskToWorker(task)
				if err != nil {
					log.Println("failed to dispatch task:", task.ID)
					continue
				}

				task.Status = "running"
				err = models.UpdateTaskStatus(task)
				if err != nil {
					log.Println("Error update task status:", err)
				}
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
