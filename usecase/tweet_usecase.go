package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type TweetUsecase interface {
	ComposeTweet(req *request.ComposeTweetRequest) (*response.TweetResponse, error)
	DeleteTweet(req *request.DeleteTweetRequest) error
}

type tweetUsecase struct {
	responser       response.TweetUsecaseResponser
	tweetRepository repository.TweetRepository
	userRepository  repository.UserRepository
}

func NewTweetUsecase(
	responser response.TweetUsecaseResponser,
	tweetRepo repository.TweetRepository,
	userRepo repository.UserRepository,
) TweetUsecase {
	return &tweetUsecase{
		responser:       responser,
		tweetRepository: tweetRepo,
		userRepository:  userRepo,
	}
}

func (u tweetUsecase) ComposeTweet(req *request.ComposeTweetRequest) (*response.TweetResponse, error) {
	user, err := u.userRepository.Find(req.UserID)
	if err != nil {
		return nil, newNoSuchUserError()
	}

	tweet := user.ComposeTweet(req.Content)
	if err := u.tweetRepository.Save(tweet); err != nil {
		return nil, fmt.Errorf("failed to compose tweet: %s", err)
	}

	return u.responser.ResponseTweet(tweet), nil
}

func (u tweetUsecase) DeleteTweet(req *request.DeleteTweetRequest) error {
	if err := u.tweetRepository.Delete(req.TweetID); err != nil {
		return fmt.Errorf("failed to delete tweet: %s", err)
	}

	return nil
}
