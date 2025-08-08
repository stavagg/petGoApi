package mocks

import (
	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stretchr/testify/mock"
)

type TodoServiceMock struct {
	mock.Mock
}

func (m *TodoServiceMock) CreateTodo(req model.CreateTodoRequest) (*model.Todo, error) {
	args := m.Called(req)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoServiceMock) GetAllTodos() ([]model.Todo, error) {
	args := m.Called()
	return args.Get(0).([]model.Todo), args.Error(1)
}

func (m *TodoServiceMock) GetTodoByID(id uint) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoServiceMock) UpdateTodo(id uint, req model.UpdateTodoRequest) (*model.Todo, error) {
	args := m.Called(id, req)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoServiceMock) DeleteTodo(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TodoServiceMock) GetTodosByCompleted(completed bool) ([]model.Todo, error) {
	args := m.Called(completed)
	return args.Get(0).([]model.Todo), args.Error(1)
}

func (m *TodoServiceMock) GetStats() (map[string]interface{}, error) {
	args := m.Called()
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (m *TodoServiceMock) ToggleTodo(id uint) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoServiceMock) MarkAllCompleted() error {
	args := m.Called()
	return args.Error(0)
}

func (m *TodoServiceMock) DeleteCompleted() error {
	args := m.Called()
	return args.Error(0)
}
