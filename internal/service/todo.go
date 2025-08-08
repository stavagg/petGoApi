package service

import (
	"errors"

	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stavagg/petGoApi/internal/repository"
)

type TodoServiceInterface interface {
	CreateTodo(req model.CreateTodoRequest) (*model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
	GetTodoByID(id uint) (*model.Todo, error)
	UpdateTodo(id uint, req model.UpdateTodoRequest) (*model.Todo, error)
	DeleteTodo(id uint) error
	GetTodosByCompleted(completed bool) ([]model.Todo, error)
	GetStats() (map[string]interface{}, error)
	ToggleTodo(id uint) (*model.Todo, error)
	MarkAllCompleted() error
	DeleteCompleted() error
}

type TodoService struct {
	repo repository.TodoRepositoryInterface
}

func NewTodoService(repo repository.TodoRepositoryInterface) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(req model.CreateTodoRequest) (*model.Todo, error) {

	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	if len(req.Title) > 255 {
		return nil, errors.New("title too long (max 255 characters)")
	}

	if len(req.Description) > 1000 {
		return nil, errors.New("description too long (max 1000 characters)")
	}

	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	err := s.repo.Create(todo)
	if err != nil {
		return nil, errors.New("failed to create todo: " + err.Error())
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("failed to get todos: " + err.Error())
	}
	return todos, nil
}

func (s *TodoService) GetTodoByID(id uint) (*model.Todo, error) {
	if id == 0 {
		return nil, errors.New("invalid todo ID")
	}

	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo(id uint, req model.UpdateTodoRequest) (*model.Todo, error) {

	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	if req.Title != "" {
		if len(req.Title) > 255 {
			return nil, errors.New("title too long (max 255 characters)")
		}
		todo.Title = req.Title
	}

	if req.Description != "" {
		if len(req.Description) > 1000 {
			return nil, errors.New("description too long (max 1000 characters)")
		}
		todo.Description = req.Description
	}

	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	err = s.repo.Update(todo)
	if err != nil {
		return nil, errors.New("failed to update todo: " + err.Error())
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(id uint) error {
	if id == 0 {
		return errors.New("invalid todo ID")
	}

	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("todo not found")
	}

	err = s.repo.Delete(id)
	if err != nil {
		return errors.New("failed to delete todo: " + err.Error())
	}

	return nil
}

func (s *TodoService) GetTodosByCompleted(completed bool) ([]model.Todo, error) {
	todos, err := s.repo.GetByCompleted(completed)
	if err != nil {
		return nil, errors.New("failed to get todos by status: " + err.Error())
	}
	return todos, nil
}

func (s *TodoService) GetStats() (map[string]interface{}, error) {
	allTodos, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("failed to get statistics: " + err.Error())
	}

	completedCount := 0
	pendingCount := 0

	for _, todo := range allTodos {
		if todo.Completed {
			completedCount++
		} else {
			pendingCount++
		}
	}

	completionRate := 0.0
	if len(allTodos) > 0 {
		completionRate = float64(completedCount) / float64(len(allTodos)) * 100
	}

	stats := map[string]interface{}{
		"total":           len(allTodos),
		"completed":       completedCount,
		"pending":         pendingCount,
		"completion_rate": completionRate,
	}

	return stats, nil
}

func (s *TodoService) ToggleTodo(id uint) (*model.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	todo.Completed = !todo.Completed

	err = s.repo.Update(todo)
	if err != nil {
		return nil, errors.New("failed to toggle todo: " + err.Error())
	}

	return todo, nil
}

func (s *TodoService) MarkAllCompleted() error {
	todos, err := s.repo.GetByCompleted(false)
	if err != nil {
		return errors.New("failed to get pending todos: " + err.Error())
	}

	for _, todo := range todos {
		todo.Completed = true
		err := s.repo.Update(&todo)
		if err != nil {
			return errors.New("failed to mark todo as completed: " + err.Error())
		}
	}

	return nil
}

func (s *TodoService) DeleteCompleted() error {
	completedTodos, err := s.repo.GetByCompleted(true)
	if err != nil {
		return errors.New("failed to get completed todos: " + err.Error())
	}

	for _, todo := range completedTodos {
		err := s.repo.Delete(todo.ID)
		if err != nil {
			return errors.New("failed to delete completed todo: " + err.Error())
		}
	}

	return nil
}
