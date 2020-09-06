package handler

import (
	"net/http"
	"todo-rest-api/internal/app/store/sqlstore"

	"github.com/gin-gonic/gin"
)

// TodosGET ...
func TodosGET(todoRepo *sqlstore.TodoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := todoRepo.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, todos)
	}
}
