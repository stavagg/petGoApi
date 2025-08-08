package model

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" binding:"required" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}
