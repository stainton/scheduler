package models

import (
	"time"

	"github.com/stainton/scheduler/internal/database"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

type Task struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

func CreateTask(task *Task) error {
	if db == nil {
		db = database.GetDB()
		db.AutoMigrate(task)
	}
	return db.Create(task).Error
}

func GetPendingTasks() ([]*Task, error) {
	if db == nil {
		db = database.GetDB()
		db.AutoMigrate(&Task{})
	}
	var tasks []*Task
	err := db.Where("status = ?", "pending").Find(&tasks).Error
	return tasks, err
}

func GetRunningTasks() ([]*Task, error) {
	if db == nil {
		db = database.GetDB()
		db.AutoMigrate(&Task{})
	}
	var tasks []*Task
	err := db.Where("status = ?", "running").Find(&tasks).Error
	return tasks, err
}

func GetTaskByID(id uint) ([]*Task, error) {
	if db == nil {
		db = database.GetDB()
		db.AutoMigrate(&Task{})
	}
	var tasks []*Task
	err := db.Where("id = ?", id).Find(&tasks).Error
	return tasks, err
}

func UpdateTaskStatus(task *Task) error {
	if db == nil {
		db = database.GetDB()
		db.AutoMigrate(task)
	}
	return db.Model(task).Update("status", task.Status).Error
}
