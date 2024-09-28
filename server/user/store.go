package user

import (
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}

//func (store *Store) GetUserByEmail(email string) (*types.User, error) {
//	// Sử dụng QueryRow thay vì Exec để lấy dữ liệu
//	row := store.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
//
//	user := new(types.User)
//
//	// Quét dữ liệu từ row vào user
//	err := row.Scan(&user.ID, &user.Name, &user.Email) // Điều chỉnh các trường dựa trên struct User
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, fmt.Errorf("user not found")
//		}
//		return nil, err
//	}
//
//	return user, nil
//}
//
//func (store *Store) GetUserByID(id string) (*types.User, error) {
//	rows, err := store.db.Query("SELECT * FROM users WHERE id=?", id)
//	if err != nil {
//		return nil, err
//	}
//	user := new(types.User)
//	for rows.Next() {
//		user, err = scanRowsIntoUser(rows)
//		if err != nil {
//			return nil, err
//		}
//	}
//	if user.ID == 0 {
//		return nil, fmt.Errorf("user not found")
//	}
//	return user, nil
//}
//func (store *Store) CreateUser(user types.User) error {
//	_, err := store.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
//	user := new(types.User)
//	err := rows.Scan(
//		&user.ID,
//		&user.FirstName,
//		&user.LastName,
//		&user.Email,
//		&user.Password)
//	if err != nil {
//		return nil, err
//	}
//	return user, nil
//}
