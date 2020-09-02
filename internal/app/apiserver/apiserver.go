package apiserver

import (
	"database/sql"
	"todo-rest-api/internal/app/handler"
	"todo-rest-api/internal/app/store/sqlstore"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	todoRepo := &sqlstore.TodoRepository{
		Store: store,
	}

	r.GET("/ping", handler.PingGET())
	r.GET("/todos", handler.TodosGET(todoRepo))
	r.POST("/todos", handler.TodosPOST(todoRepo))

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
