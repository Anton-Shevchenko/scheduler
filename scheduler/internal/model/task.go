package model

type Task struct {
	ID       int    `json:"id"`
	JobName  string `json:"job_name" binding:"required"`
	TimeToDo int64  `json:"time_to_do" binding:"required"`
	Status   string `json:"status" default:"todo"`
}
