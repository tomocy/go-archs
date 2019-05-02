package model

type UserID string

type User struct {
	ID               UserID
	Email            string
	isPasswordHashed bool
	Password         string
}

func (u *User) ValidateSelf() error {
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

	return nil
}
