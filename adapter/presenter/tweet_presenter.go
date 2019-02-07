package presenter

import (
	"github.com/tomocy/archs/domain/model"
)

type TweetPresenter interface {
	PresentTweet(tweet *model.Tweet) *TweetPresent
}

func NewTweetPresenter() TweetPresenter {
	return new(tweetPresenter)
}

type tweetPresenter struct {
}

type TweetPresent struct {
	ID      string
	UserID  string
	Content string
}

func (p tweetPresenter) PresentTweet(tweet *model.Tweet) *TweetPresent {
	return &TweetPresent{
		ID:      string(tweet.ID),
		UserID:  string(tweet.UserID),
		Content: tweet.Content,
	}
}
