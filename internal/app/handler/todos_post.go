package handler

import (
	"fmt"
	"net/http"
	"todo-rest-api/internal/app/model"
	"todo-rest-api/internal/app/store/sqlstore"

	"github.com/gin-gonic/gin"
)

type todosPostRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// TodosPOST ...
func TodosPOST(todoRepo *sqlstore.TodoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := todosPostRequest{}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newTodo := model.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		newTodo.BeforeCreate()

		todoRepo.Create(&newTodo)

		c.Status(http.StatusNoContent)
	}
}
