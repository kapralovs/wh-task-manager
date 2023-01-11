package models

import "time"

type Product struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Название
	Name string `json:"name,omitempty"`
	// Количество
	Count int64 `json:"count,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
