package main

import (
	"todo-rest-api/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGET())

	r.Run()
}
