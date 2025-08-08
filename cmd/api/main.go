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
	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r := gin.Default()

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

	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.GetAllTodos)
			todos.GET("/stats", todoHandler.GetStats)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
			todos.POST("/:id/toggle", todoHandler.ToggleTodo)
		}
	}

	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Printf("🔗 API documentation: http://localhost%s/", cfg.Port)
	log.Printf("📝 API endpoints: http://localhost%s/api/v1/todos", cfg.Port)
	log.Fatal(r.Run(cfg.Port))
}
