package usecase

import (
	"github.com/kapralovs/wh-task-manager/internal/models"
	"github.com/kapralovs/wh-task-manager/internal/task"
)

type TaskUsecase struct {
}

func NewTaskUsecase(tasksRepo task.Repository) *TaskUsecase {
	return &TaskUsecase{}
}

func (uc *TaskUsecase) CreateTask(t *models.Task) error { return nil }

func (uc *TaskUsecase) UpdateTask(t *models.Task, id int64) error { return nil }

func (uc *TaskUsecase) DeleteTask(id int64) error { return nil }

func (uc *TaskUsecase) GetTask(id int64) (*models.Task, error) {
	var t *models.Task
	return t, nil
}

func (uc *TaskUsecase) GetTasksList() ([]*models.Task, error) {
	tasksList := []*models.Task{}
	return tasksList, nil
}
