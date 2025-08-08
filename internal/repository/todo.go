package repository

import (
	"github.com/stavagg/petGoApi/internal/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// Create создает новую задачу
func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

// GetAll получает все задачи
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Order("created_at desc").Find(&todos).Error
	return todos, err
}

// GetByID получает задачу по ID
func (r *TodoRepository) GetByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Update обновляет задачу
func (r *TodoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

// Delete удаляет задачу по ID
func (r *TodoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}

// GetByCompleted получает задачи по статусу выполнения
func (r *TodoRepository) GetByCompleted(completed bool) ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Where("completed = ?", completed).Order("created_at desc").Find(&todos).Error
	return todos, err
}
