package repository

import (
	"fmt"
	"gorm.io/gorm"
	"taskManager/internal/interfaces"
	"taskManager/internal/model"
	"time"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) interfaces.TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *TaskRepository) GetAllToDo(duration time.Duration) ([]model.Task, error) {
	var tasks []model.Task

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, nil
	}

	now := time.Now().In(loc)
	oneDayAgo := now.Unix() - 86400

	err = r.db.
		Where("status != ?", "sent").
		Where("time_to_do > ?", oneDayAgo).
		Where("time_to_do < ?", now.Unix()).
		Find(&tasks).
		Error

	return tasks, err
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task

	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepository) Save(task *model.Task) error {

	return r.db.Save(&task).Error
}
