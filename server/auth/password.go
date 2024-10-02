package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hass), nil
}
func CheckPassword(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
