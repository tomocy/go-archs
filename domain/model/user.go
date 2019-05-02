package model

type UserID string

type User struct {
	ID       UserID
	Email    string
	Password string

	isIDAllocated    bool
	isPasswordHashed bool
}

func (u *User) AllocateID(id UserID) error {
	if u.isIDAllocated {
		return errorf("user", "id is already allocated")
	}

	u.ID = id
	u.isIDAllocated = true

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

	return nil
}
