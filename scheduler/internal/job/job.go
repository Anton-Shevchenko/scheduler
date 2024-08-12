package job

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"taskManager/config"
	"taskManager/internal/interfaces"
	"time"
)

type Job struct {
	taskRepo interfaces.TaskRepository
	Cnf      *config.Config
}

func NewJob(taskRepo interfaces.TaskRepository, cnf *config.Config) *Job {
	return &Job{
		taskRepo: taskRepo,
		Cnf:      cnf,
	}
}

func (j *Job) Do() {
	tasks, err := j.taskRepo.GetAllToDo(time.Hour * 24)

	if err != nil {
		fmt.Println("Task err", err.Error())
		return
	}

	for _, task := range tasks {
		j.sendMessage(task.JobName)
		task.Status = "sent"
		err := j.taskRepo.Save(&task)
		if err != nil {
			fmt.Println("Task err", err.Error())
			return
		}
	}
}

func (j *Job) sendMessage(text string) {
	baseURL := "https://api.telegram.org/bot"
	token := j.Cnf.TGToken
	chatID := j.Cnf.TGChatId
	endpoint := fmt.Sprintf("%s%s/sendMessage", baseURL, token)
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", text)

	response, err := http.PostForm(endpoint, data)
	if err != nil {
		fmt.Printf("Error sending message: %s\n", err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err.Error())
		return
	}
	fmt.Printf("Response from Telegram: %s\n", body)
}
