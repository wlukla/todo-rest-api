package handler

import (
	"net/http"
	"todo-rest-api/platform/todo"

	"github.com/gin-gonic/gin"
)

// TodosGET ...
func TodosGET(todos *todo.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, todos.GetAll())
	}
}
