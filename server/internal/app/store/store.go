package store

// Store ...
type Store interface {
	TodoRepository() *TodoRepository
}
