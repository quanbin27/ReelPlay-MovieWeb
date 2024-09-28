package types

import "time"

type UserStore interface {
	//GetUserByEmail(email string) (*User, error)
}

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
