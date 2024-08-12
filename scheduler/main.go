package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"taskManager/config"
	"taskManager/internal/db"
	"taskManager/internal/http/handler"
	api_router "taskManager/internal/http/v1/router"
	"taskManager/internal/job"
	"taskManager/internal/model"
	"taskManager/internal/repository"
	"taskManager/pkg/jobManager"
	"time"
)

// should be in /cmd/app
func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error: Load Config: %v", err)
	}

	postgres := db.NewDbConn()
	migrations(postgres)

	taskRepo := repository.NewTaskRepository(postgres)
	taskHandler := handler.NewTaskHandler(taskRepo)

	v1 := api_router.NewRouter(taskHandler)

	jobber := jobManager.NewJobManager()
	//Jobs
	jobService := job.NewJob(taskRepo, cnf)

	jobber.Add(jobService.Do, "job", 1*time.Minute)
	jobber.Run()

	v1.Serve(cnf)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.Task{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
