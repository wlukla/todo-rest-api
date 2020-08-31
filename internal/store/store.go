package store

type Store interface {
	TodoRepository() *TodoRepository
}
