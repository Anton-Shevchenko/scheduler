package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskManager/internal/interfaces"
	"taskManager/internal/model"
)

type TaskHandler struct {
	repo interfaces.TaskRepository
}

func NewTaskHandler(repo interfaces.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

func (h *TaskHandler) Add(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = "todo"

	createdTask, err := h.repo.Create(&task)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": createdTask,
	})
}

func (h *TaskHandler) Get(c *gin.Context) {
	tasks, err := h.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
