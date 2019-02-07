package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/usecase/request"
)

type TweetUsecase interface {
	ComposeTweet(req *request.ComposeTweetRequest) (*model.Tweet, error)
	DeleteTweet(req *request.DeleteTweetRequest) error
}

type tweetUsecase struct {
	tweetRepository repository.TweetRepository
	userRepository  repository.UserRepository
}

func NewTweetUsecase(
	tweetRepo repository.TweetRepository,
	userRepo repository.UserRepository,
) TweetUsecase {
	return &tweetUsecase{
		tweetRepository: tweetRepo,
		userRepository:  userRepo,
	}
}

func (u tweetUsecase) ComposeTweet(req *request.ComposeTweetRequest) (*model.Tweet, error) {
	user, err := u.userRepository.Find(req.UserID)
	if err != nil {
		return nil, newNoSuchUserError()
	}

	tweet := user.ComposeTweet(req.Content)
	if err := u.tweetRepository.Save(tweet); err != nil {
		return nil, fmt.Errorf("failed to compose tweet: %s", err)
	}

	return tweet, nil
}

func (u tweetUsecase) DeleteTweet(req *request.DeleteTweetRequest) error {
	if err := u.tweetRepository.Delete(req.TweetID); err != nil {
		return fmt.Errorf("failed to delete tweet: %s", err)
	}

	return nil
}
