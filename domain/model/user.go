package model

type UserID string

type User struct {
	ID               UserID
	Email            string
	isPasswordHashed bool
	Password         string
}
