package handler

import (
	"net/http"
	"todo-rest-api/platform/todo"

	"github.com/gin-gonic/gin"
)

type todosPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodosPOST ...
func TodosPOST(todos todo.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := todosPostRequest{}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Bind(requestBody)

		newTodo := todo.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}

		todos.Add(newTodo)

		c.Status(http.StatusNoContent)
	}
}
