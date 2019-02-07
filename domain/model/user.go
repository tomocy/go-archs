package model

type UserID string

type User struct {
	ID       UserID
	Email    string
	Password string
}

func NewUser(id UserID, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

func (u User) ComposeTweet(id TweetID, content string) *Tweet {
	return NewTweet(id, u.ID, content)
}
