package sqlstore

import "todo-rest-api/internal/app/model"

//UserRepository ...
type UserRepository struct {
	Store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := r.Store.db.QueryRow(
		"INSERT INTO users (username, password) values ($1, $2) RETURNING id",
		u.Username, u.Password,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}
