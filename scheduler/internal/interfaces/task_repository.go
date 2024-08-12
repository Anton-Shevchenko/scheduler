package interfaces

import (
	"taskManager/internal/model"
	"time"
)

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	GetAllToDo(duration time.Duration) ([]model.Task, error)
	Save(task *model.Task) error
	GetAll() ([]model.Task, error)
}
