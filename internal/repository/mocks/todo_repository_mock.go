package mocks

import (
	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) Create(todo *model.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *TodoRepositoryMock) GetAll() ([]model.Todo, error) {
	args := m.Called()
	return args.Get(0).([]model.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) GetByID(id uint) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Update(todo *model.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *TodoRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TodoRepositoryMock) GetByCompleted(completed bool) ([]model.Todo, error) {
	args := m.Called(completed)
	return args.Get(0).([]model.Todo), args.Error(1)
}
