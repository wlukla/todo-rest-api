package handler

import (
	"net/http"
	"todo-rest-api/internal/app/model"
	"todo-rest-api/internal/app/store/sqlstore"

	"github.com/gin-gonic/gin"
)

type usersPostRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UsersPOST ...
func UsersPOST(userRepo *sqlstore.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := usersPostRequest{}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser := model.User{
			Username: request.Username,
			Password: request.Password,
		}

		if err := userRepo.Create(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, newUser)
	}
}
