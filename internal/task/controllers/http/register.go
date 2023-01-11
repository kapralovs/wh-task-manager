package http

import (
	"github.com/kapralovs/wh-task-manager/internal/task"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(r *echo.Group, uc task.Usecase) {
	h := NewHandler(uc)
	r.POST("create", h.createTask)
	r.PUT("update/:id", h.updateTask)
	r.PUT("delete/:id", h.deleteTask)
	r.GET("get/:id", h.getTask)
	r.GET("get-list/", h.getTasksList)
}
