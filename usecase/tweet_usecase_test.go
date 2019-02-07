package usecase

import (
	"testing"

	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/usecase/request"
)

func TestComposeTweet(t *testing.T) {
	tweetRepo := memory.NewTweetRepository()
	userRepo := memory.NewUserRepository()
	usecase := NewTweetUsecase(presenter.NewTweetUsecaseResponser(), tweetRepo, userRepo)
	userID := model.UserID("test user id")
	content := "Is this a pen?"
	userRepo.Save(&model.User{
		ID: userID,
	})
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{
			"normal",
			func(t *testing.T) {
				req := request.NewComposeTweetRequest(userID, content)
				_, err := usecase.ComposeTweet(req)
				if err != nil {
					t.Errorf("unexpected error: %s\n", err)
				}
			},
		},
		{
			"no such user",
			func(t *testing.T) {
				req := request.NewComposeTweetRequest(model.UserID(""), content)
				_, err := usecase.ComposeTweet(req)
				if !IsNoSuchUserError(err) {
					t.Errorf("unexpected error: got %s, but expected NoSuchUserError\n", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.tester)
	}
}

func TestDeleteTweet(t *testing.T) {
	tweetRepo := memory.NewTweetRepository()
	userRepo := memory.NewUserRepository()
	usecase := NewTweetUsecase(presenter.NewTweetUsecaseResponser(), tweetRepo, userRepo)
	tweetID := model.TweetID("test tweet id")
	tweetRepo.Save(&model.Tweet{
		ID: tweetID,
	})
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{
			"normal",
			func(t *testing.T) {
				req := request.NewDeleteTweetRequest(tweetID)
				if err := usecase.DeleteTweet(req); err != nil {
					t.Errorf("unexpected error: %s\n", err)
				}
			},
		},
		{
			"no such tweet",
			func(t *testing.T) {
				req := request.NewDeleteTweetRequest("")
				if err := usecase.DeleteTweet(req); err != nil {
					t.Errorf("unexpected error: %s\n", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.tester)
	}
}
