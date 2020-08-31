package apiserver

import (
	"database/sql"
	"todo-rest-api/internal/handler"
	sqlstore "todo-rest-api/internal/store"
	"todo-rest-api/platform/todo"

	"github.com/gin-gonic/gin"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	r := gin.Default()
	todoRepo := todo.New()

	r.GET("/ping", handler.PingGET())
	r.GET("/todos", handler.TodosGET(todos))
	r.POST("/todos", handler.TodosPOST(todos))

	r.Run()

	return nil
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
