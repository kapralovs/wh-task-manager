package models

type Task struct {
	ID                    int64
	Name                  string
	CreatedBy             int64
	IsAutomaticalyCreated bool
	IsInternal            bool
	Content               []*Product
}
