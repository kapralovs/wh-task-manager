package usecase

import (
	"github.com/kapralovs/wh-task-manager/internal/models"
	"github.com/kapralovs/wh-task-manager/internal/task"
)

type TaskUsecase struct {
	Repo task.Repository
}

func NewTaskUsecase(tasksRepo task.Repository) *TaskUsecase {
	return &TaskUsecase{
		Repo: tasksRepo,
	}
}

func (uc *TaskUsecase) CreateTask(t *models.Task) error {
	err := uc.Repo.CreateTask(t)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TaskUsecase) UpdateTask(t *models.Task, id int64) error {
	err := uc.Repo.UpdateTask(t, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TaskUsecase) DeleteTask(id int64) error {
	err := uc.Repo.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TaskUsecase) GetTask(id int64) (*models.Task, error) {
	t, err := uc.Repo.GetTask(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (uc *TaskUsecase) GetTasksList() ([]*models.Task, error) {
	tList, err := uc.Repo.GetTasksList()
	if err != nil {
		return nil, err
	}

	return tList, nil
}
