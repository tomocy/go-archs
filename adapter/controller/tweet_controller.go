package controller

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type TweetController interface {
	ComposeTweet(userID, content string) (*model.Tweet, error)
}

type tweetController struct {
	usecase usecase.TweetUsecase
}

func NewTweetController(usecase usecase.TweetUsecase) TweetController {
	return &tweetController{
		usecase: usecase,
	}
}

func (c tweetController) ComposeTweet(userID, content string) (*model.Tweet, error) {
	return c.usecase.ComposeTweet(
		request.NewComposeTweetRequest(model.UserID(userID), content),
	)
}
