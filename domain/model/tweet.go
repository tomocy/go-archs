package model

import (
	"time"
)

type TweetID string

type Tweet struct {
	ID        TweetID
	UserID    UserID
	Content   string
	CreatedAt time.Time
}

func NewTweet(id TweetID, userID UserID, content string) *Tweet {
	return &Tweet{
		ID:        id,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
