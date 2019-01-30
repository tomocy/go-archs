package model

import (
	"time"

	"github.com/google/uuid"
)

type TweetID string

func generateTweetID() TweetID {
	return TweetID(uuid.New().String())
}

type Tweet struct {
	ID        TweetID
	UserID    UserID
	Content   string
	CreatedAt time.Time
}

func NewTweet(userID UserID, content string) *Tweet {
	return &Tweet{
		ID:        generateTweetID(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
