package model

import "github.com/tomocy/archs/domain/service"

type UserID string

type User struct {
	ID       UserID
	Email    string
	Password string

	isIDAllocated    bool
	isPasswordHashed bool
	isValidated      bool
}

func (u *User) AllocateID(id UserID) error {
	if u.isIDAllocated {
		return errorf("user", "id is already allocated")
	}

	u.ID = id
	u.isIDAllocated = true

	return nil
}

func (u *User) HashPassword(service service.HashService) error {
	plain := u.Password
	if isEmpty(plain) {
		return errorf("user", "password is empty")
	}
	hash, err := service.Hash(plain)
	if err != nil {
		return err
	}

	u.Password = hash
	u.isPasswordHashed = true

	return nil
}

func (u *User) ValidateSelf() error {
	if !u.isIDAllocated {
		return errorf("user", "id is not allocated")
	}
	if isEmpty(string(u.ID)) {
		return errorf("user", "id is empty")
	}
	if isEmpty(u.Email) {
		return errorf("user", "email is empty")
	}
	if isEmpty(u.Password) {
		return errorf("user", "password is empty")
	}
	if !u.isPasswordHashed {
		return errorf("user", "password is not hashed")
	}

	u.isValidated = true

	return nil
}

func (u *User) IsValidated() bool {
	return u.isValidated
}
