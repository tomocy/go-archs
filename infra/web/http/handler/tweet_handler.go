package handler

import (
	"fmt"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/session"
	"github.com/tomocy/archs/infra/web/http/validator"
	"github.com/tomocy/archs/usecase"
)

type TweetHandler interface {
	ComposeTweet(w http.ResponseWriter, r *http.Request)
}

type tweetHandler struct {
	controller controller.TweetController
}

func NewTweetHandler(controller controller.TweetController) TweetHandler {
	return &tweetHandler{
		controller: controller,
	}
}

func (h tweetHandler) ComposeTweet(w http.ResponseWriter, r *http.Request) {
	userID := session.SessionService.GetAuthenticUserID(r)
	validated, err := validator.ValidateToComposeTweet(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tweet, err := h.controller.ComposeTweet(userID, validated.Content)
	if err != nil {
		switch err.(type) {
		case usecase.NoSuchUserError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "compose tweet: {ID: %s, UserID: %s, Content: %s}\n", tweet.ID, tweet.UserID, tweet.Content)
}