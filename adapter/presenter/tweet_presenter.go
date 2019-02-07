package presenter

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase/response"
)

func NewTweetUsecaseResponser() response.TweetUsecaseResponser {
	return new(tweetPresenter)
}

type tweetPresenter struct {
}

func (p tweetPresenter) ResponseTweet(tweet *model.Tweet) *response.TweetResponse {
	return response.NewTweetResponse(string(tweet.ID), string(tweet.UserID), tweet.Content)
}
