package api_router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taskManager/config"
	"taskManager/internal/http/handler"
)

type Router struct {
	taskHandler *handler.TaskHandler
}

func NewRouter(taskHandler *handler.TaskHandler) *Router {
	return &Router{taskHandler: taskHandler}
}

func (r *Router) Serve(cnf *config.Config) {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")
	taskApi := api.Group("/task")

	taskApi.POST("/", r.taskHandler.Add)
	taskApi.GET("/", r.taskHandler.Get)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
