package handler_test

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stavagg/petGoApi/internal/handler"
	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stavagg/petGoApi/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodo_Handler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	serviceMock := new(mocks.TodoServiceMock)
	h := handler.NewTodoHandler(serviceMock)

	todo := &model.Todo{ID: 1, Title: "Test", Description: "Desc"}
	serviceMock.On("CreateTodo", mock.AnythingOfType("model.CreateTodoRequest")).Return(todo, nil)

	reqBody := `{"title":"Test","description":"Desc"}`
	req := httptest.NewRequest("POST", "/todos", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	h.CreateTodo(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	serviceMock.AssertExpectations(t)
}

func TestGetAllTodos_Handler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	serviceMock := new(mocks.TodoServiceMock)
	h := handler.NewTodoHandler(serviceMock)

	todos := []model.Todo{{ID: 1, Title: "Test"}}
	serviceMock.On("GetAllTodos").Return(todos, nil)

	req := httptest.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	h.GetAllTodos(c)

	assert.Equal(t, http.StatusOK, w.Code)
	serviceMock.AssertExpectations(t)
}
