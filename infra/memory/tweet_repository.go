package memory

import (
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
)

var TweetRepository repository.TweetRepository = NewTweetRepository()

func NewTweetRepository() repository.TweetRepository {
	return newTweetRepository()
}

type tweetRepository struct {
	tweets []*model.Tweet
}

func newTweetRepository() *tweetRepository {
	return new(tweetRepository)
}

func (r *tweetRepository) Save(tweet *model.Tweet) error {
	r.tweets = append(r.tweets, tweet)
	return nil
}
