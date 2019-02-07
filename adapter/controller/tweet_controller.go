package controller

import (
	"net/http"

	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/request"
)

type TweetController interface {
	GetAuthenticUserID(r *http.Request) string
	ComposeTweet(userID, content string) (*presenter.TweetPresent, error)
	DeleteTweet(tweetID string) error
}

type tweetController struct {
	AuthenticationController
	presenter presenter.TweetPresenter
	usecase   usecase.TweetUsecase
}

func NewTweetController(
	authController AuthenticationController,
	presenter presenter.TweetPresenter,
	usecase usecase.TweetUsecase,
) TweetController {
	return &tweetController{
		AuthenticationController: authController,
		presenter:                presenter,
		usecase:                  usecase,
	}
}

func (c tweetController) ComposeTweet(userID, content string) (*presenter.TweetPresent, error) {
	tweet, err := c.usecase.ComposeTweet(
		request.NewComposeTweetRequest(model.UserID(userID), content),
	)
	return c.presenter.PresentTweet(tweet), err
}

func (c tweetController) DeleteTweet(tweetID string) error {
	return c.usecase.DeleteTweet(
		request.NewDeleteTweetRequest(model.TweetID(tweetID)),
	)
}
