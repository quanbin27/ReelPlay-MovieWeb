package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
