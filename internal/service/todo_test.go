package service_test

import (
	"errors"
	"testing"

	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stavagg/petGoApi/internal/repository/mocks"
	"github.com/stavagg/petGoApi/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodo_Success(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	req := model.CreateTodoRequest{Title: "Test", Description: "Desc"}
	repoMock.On("Create", mock.AnythingOfType("*model.Todo")).Return(nil)

	todo, err := svc.CreateTodo(req)
	assert.NoError(t, err)
	assert.Equal(t, "Test", todo.Title)
	assert.Equal(t, "Desc", todo.Description)

	repoMock.AssertExpectations(t)
}

func TestCreateTodo_EmptyTitle(t *testing.T) {
	svc := service.NewTodoService(nil)

	_, err := svc.CreateTodo(model.CreateTodoRequest{Title: "", Description: "desc"})
	assert.EqualError(t, err, "title is required")
}

func TestCreateTodo_RepoError(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	repoMock.On("Create", mock.Anything).Return(errors.New("db error"))

	_, err := svc.CreateTodo(model.CreateTodoRequest{Title: "T", Description: ""})
	assert.ErrorContains(t, err, "failed to create todo")
	repoMock.AssertExpectations(t)
}

func TestGetAllTodos_Success(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	sample := []model.Todo{
		{ID: 1, Title: "A", Description: "a", Completed: false},
		{ID: 2, Title: "B", Description: "b", Completed: true},
	}
	repoMock.On("GetAll").Return(sample, nil)

	todos, err := svc.GetAllTodos()
	assert.NoError(t, err)
	assert.Equal(t, sample, todos)

	repoMock.AssertExpectations(t)
}

func TestGetTodoByID_NotFound(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	repoMock.On("GetByID", uint(1)).Return((*model.Todo)(nil), errors.New("not found"))

	_, err := svc.GetTodoByID(1)
	assert.EqualError(t, err, "todo not found")
	repoMock.AssertExpectations(t)
}

func TestUpdateTodo_Success(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	existing := &model.Todo{ID: 1, Title: "Old", Description: "old", Completed: false}
	repoMock.On("GetByID", uint(1)).Return(existing, nil)
	repoMock.On("Update", existing).Return(nil)

	req := model.UpdateTodoRequest{Title: "New", Description: "new", Completed: new(bool)}
	*req.Completed = true

	updated, err := svc.UpdateTodo(1, req)
	assert.NoError(t, err)
	assert.Equal(t, "New", updated.Title)
	assert.True(t, updated.Completed)

	repoMock.AssertExpectations(t)
}

func TestDeleteTodo_Success(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	repoMock.On("GetByID", uint(1)).Return(&model.Todo{ID: 1}, nil)
	repoMock.On("Delete", uint(1)).Return(nil)

	err := svc.DeleteTodo(1)
	assert.NoError(t, err)
	repoMock.AssertExpectations(t)
}

func TestGetStats_Calculation(t *testing.T) {
	repoMock := new(mocks.TodoRepositoryMock)
	svc := service.NewTodoService(repoMock)

	sample := []model.Todo{
		{Completed: true},
		{Completed: false},
		{Completed: true},
	}
	repoMock.On("GetAll").Return(sample, nil)

	stats, err := svc.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, 3, stats["total"])
	assert.Equal(t, 2, stats["completed"])
	assert.Equal(t, 1, stats["pending"])
	assert.InDelta(t, 66.666, stats["completion_rate"], 0.1)

	repoMock.AssertExpectations(t)
}
