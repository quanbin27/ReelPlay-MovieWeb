package user

import (
	"database/sql"
	"fmt"
	"github.com/quanbin27/ReelPlay/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}
func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
