package handler

import (
	"net/http"
	"todo-rest-api/internal/app/model"
	"todo-rest-api/internal/app/store/sqlstore"

	"github.com/gin-gonic/gin"
)

type todosPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodosPOST ...
func TodosPOST(todoRepo *sqlstore.TodoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := todosPostRequest{}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Bind(requestBody)

		newTodo := model.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		newTodo.BeforeCreate()

		todoRepo.Create(&newTodo)

		c.Status(http.StatusNoContent)
	}
}
