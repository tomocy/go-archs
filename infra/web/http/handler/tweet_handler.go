package handler

import (
	"fmt"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/web/http/validator"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/chi"
)

type TweetHandler interface {
	ComposeTweet(w http.ResponseWriter, r *http.Request)
	DeleteTweet(w http.ResponseWriter, r *http.Request)
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
	userID := h.controller.GetAuthenticUserID(r)
	validated, err := validator.ValidateToComposeTweet(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.controller.ComposeTweet(userID, validated.Content)
	if err != nil {
		switch err.(type) {
		case usecase.NoSuchUserError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "compose resp: %v\n", resp)
}

func (h tweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := chi.URLParam(r, "id")
	if err := h.controller.DeleteTweet(tweetID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "delete tweet")
}
