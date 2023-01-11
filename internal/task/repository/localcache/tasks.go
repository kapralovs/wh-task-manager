package localcache

import (
	"errors"
	"sync"

	"github.com/kapralovs/wh-task-manager/internal/models"
)

type LocalRepo struct {
	mu    *sync.Mutex
	tasks map[int64]*models.Task
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		mu:    new(sync.Mutex),
		tasks: make(map[int64]*models.Task),
	}
}

func (r *LocalRepo) CreateTask(t *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[int64(len(r.tasks)+1)] = t

	return nil
}

func (r *LocalRepo) UpdateTask(t *models.Task, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[int64(len(r.tasks)+1)] = t

	return nil
}

func (r *LocalRepo) DeleteTask(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.tasks, id)

	return nil
}

func (r *LocalRepo) GetTask(id int64) (*models.Task, error) {
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task does not exist")
	}

	return t, nil
}

func (r *LocalRepo) GetTasksList() ([]*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var tasksList = make([]*models.Task, len(r.tasks))

	for _, t := range r.tasks {
		tasksList = append(tasksList, t)
	}

	return tasksList, nil
}
