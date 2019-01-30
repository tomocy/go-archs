package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/usecase/request"
)

func TestComposeTweet(t *testing.T) {
	userRepo := memory.NewUserRepository()
	usecase := NewTweetUsecase(userRepo)
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
