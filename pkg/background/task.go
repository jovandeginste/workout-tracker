package background

import (
	"gorm.io/gorm"
)

type Task interface {
	Run(db *gorm.DB) error
	TaskType() TaskType
}

type TaskType string
