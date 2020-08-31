package handler

import (
	"net/http"
	"todo-rest-api/internal/store"

	"github.com/gin-gonic/gin"
)

// TodosGET ...
func TodosGET(store store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, store.TodoRepository().GetAll())
	}
}
