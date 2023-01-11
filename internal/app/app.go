package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kapralovs/wh-task-manager/internal/task"
	taskhttp "github.com/kapralovs/wh-task-manager/internal/task/controllers/http"
	taskrepo "github.com/kapralovs/wh-task-manager/internal/task/repository/localcache"
	taskusecase "github.com/kapralovs/wh-task-manager/internal/task/usecase"
	"github.com/labstack/echo/v4"
)

type app struct {
	httpServer  *http.Server
	taskUsecase task.Usecase
}

func NewApp() *app {
	tasksRepo := taskrepo.NewLocalRepo()

	return &app{
		taskUsecase: taskusecase.NewTaskUsecase(tasksRepo),
	}
}

func (a *app) Run(port string) error {
	router := echo.New()
	// authGroup := router.Group("/auth/")
	// noteGroup := router.Group("/note/", authhttp.NewAuthMiddlewareHandler(a.authUseCase))
	// authhttp.RegisterEndpoints(authGroup, a.authUseCase)
	// notehttp.RegisterEndpoints(noteGroup, a.noteUseCase)

	taskGroup := router.Group("/task/")
	taskhttp.RegisterEndpoints(taskGroup, a.taskUsecase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println("Server is starting...")
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
