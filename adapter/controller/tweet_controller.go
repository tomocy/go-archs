package controller

import (
	"net/http"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type TweetController interface {
	GetAuthenticUserID(r *http.Request) string
	ComposeTweet(userID, content string) (*model.Tweet, error)
	DeleteTweet(tweetID string) error
}

type tweetController struct {
	AuthenticationController
	usecase usecase.TweetUsecase
}

func NewTweetController(authController AuthenticationController, usecase usecase.TweetUsecase) TweetController {
	return &tweetController{
		AuthenticationController: authController,
		usecase:                  usecase,
	}
}

func (c tweetController) ComposeTweet(userID, content string) (*model.Tweet, error) {
	return c.usecase.ComposeTweet(
		request.NewComposeTweetRequest(model.UserID(userID), content),
	)
}

func (c tweetController) DeleteTweet(tweetID string) error {
	return c.usecase.DeleteTweet(
		request.NewDeleteTweetRequest(model.TweetID(tweetID)),
	)
}
