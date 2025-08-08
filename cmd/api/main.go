package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/stavagg/petGoApi/internal/config"
	"github.com/stavagg/petGoApi/internal/handler"
	"github.com/stavagg/petGoApi/internal/model"
	"github.com/stavagg/petGoApi/internal/repository"
	"github.com/stavagg/petGoApi/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Подключение к базе данных
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автомиграция (создание таблиц)
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Инициализация слоев (снизу вверх)
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// Настройка Gin роутера
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Базовые роуты
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "🚀 PetGoApi is running!",
			"version": "1.0.0",
			"status":  "healthy",
			"endpoints": []string{
				"GET /api/v1/todos - получить все задачи",
				"POST /api/v1/todos - создать задачу",
				"GET /api/v1/todos/:id - получить задачу по ID",
				"PUT /api/v1/todos/:id - обновить задачу",
				"DELETE /api/v1/todos/:id - удалить задачу",
				"POST /api/v1/todos/:id/toggle - переключить статус",
				"GET /api/v1/todos/stats - статистика",
			},
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "connected"})
	})

	// API роуты v1
	api := r.Group("/api/v1")
	{
		// Todo endpoints
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)            // POST /api/v1/todos
			todos.GET("", todoHandler.GetAllTodos)            // GET /api/v1/todos
			todos.GET("/stats", todoHandler.GetStats)         // GET /api/v1/todos/stats
			todos.GET("/:id", todoHandler.GetTodoByID)        // GET /api/v1/todos/1
			todos.PUT("/:id", todoHandler.UpdateTodo)         // PUT /api/v1/todos/1
			todos.DELETE("/:id", todoHandler.DeleteTodo)      // DELETE /api/v1/todos/1
			todos.POST("/:id/toggle", todoHandler.ToggleTodo) // POST /api/v1/todos/1/toggle
		}
	}

	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Printf("🔗 API documentation: http://localhost%s/", cfg.Port)
	log.Printf("📝 API endpoints: http://localhost%s/api/v1/todos", cfg.Port)
	log.Fatal(r.Run(cfg.Port))

}
