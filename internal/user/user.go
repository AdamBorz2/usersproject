package user

import "github.com/google/uuid"

type User struct {
	UserID       uuid.UUID `json:"userId"`
	Fname        string    `json:"fname"`
	Lname        string    `json:"lname"`
	Age          int       `json:"age"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
}
