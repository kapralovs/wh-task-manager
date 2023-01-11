package task

import "github.com/kapralovs/wh-task-manager/internal/models"

type Repository interface {
	CreateTask(t *models.Task) error
	UpdateTask(t *models.Task, id int64) error
	DeleteTask(id int64) error
	GetTask(id int64) (*models.Task, error)
	GetTasksList() ([]*models.Task, error)
}
type Usecase interface {
	CreateTask(t *models.Task) error
	UpdateTask(t *models.Task, id int64) error
	DeleteTask(id int64) error
	GetTask(id int64) (*models.Task, error)
	GetTasksList() ([]*models.Task, error)
}
