package http

import (
	"net/http"
	"strconv"

	"github.com/kapralovs/wh-task-manager/internal/models"
	"github.com/kapralovs/wh-task-manager/internal/task"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	usecase task.Usecase
}

func NewHandler(uc task.Usecase) *TaskHandler {
	return &TaskHandler{usecase: uc}
}

func (h *TaskHandler) createTask(c echo.Context) error {
	var t *models.Task
	err := c.Bind(t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request!")
	}
	err = h.usecase.CreateTask(t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "task is not created!")
	}

	return c.JSON(http.StatusCreated, "task created!")
}

func (h *TaskHandler) updateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request!")
	}

	var t *models.Task
	err = c.Bind(t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request!")
	}

	err = h.usecase.UpdateTask(t, int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "task is not updated!")
	}

	return c.JSON(http.StatusOK, "task updated!")
}

func (h *TaskHandler) deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request!")
	}

	err = h.usecase.DeleteTask(int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "task is not deleted!")
	}

	return c.JSON(http.StatusOK, "task deleted!")
}

func (h *TaskHandler) getTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request!")
	}

	t, err := h.usecase.GetTask(int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, t)
	}

	return c.JSON(http.StatusOK, t)
}

func (h *TaskHandler) getTasksList(c echo.Context) error {
	tasks, err := h.usecase.GetTasksList()
	if err != nil {
		return c.JSON(http.StatusNotFound, tasks)
	}
	return c.JSON(http.StatusOK, tasks)
}
