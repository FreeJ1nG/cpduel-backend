package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	PasswordHash string `json:"password_hash"`
}

func (u *User) ValidatePasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
