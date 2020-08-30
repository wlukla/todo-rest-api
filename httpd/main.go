package main

import (
	"todo-rest-api/httpd/handler"
	"todo-rest-api/platform/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	todos := todo.New()

	r.GET("/ping", handler.PingGET())
	r.GET("/todos", handler.TodosGET(todos))
	r.POST("/todos", handler.TodosPOST(todos))

	r.Run()
}
