package usecase

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/usecase/request"
)

type TweetUsecase interface {
	ComposeTweet(req *request.ComposeTweetRequest) (*model.Tweet, error)
}

type tweetUsecase struct {
	userRepository repository.UserRepository
}

func NewTweetUsecase(userRepo repository.UserRepository) TweetUsecase {
	return &tweetUsecase{
		userRepository: userRepo,
	}
}

func (u tweetUsecase) ComposeTweet(req *request.ComposeTweetRequest) (*model.Tweet, error) {
	user, err := u.userRepository.Find(req.UserID)
	if err != nil {
		return nil, newNoSuchUserError()
	}

	return user.ComposeTweet(req.Content), nil
}
