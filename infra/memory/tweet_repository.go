package memory

import (
	"github.com/google/uuid"
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

func (r tweetRepository) NextID() model.TweetID {
	return model.TweetID(uuid.New().String())
}

func (r *tweetRepository) Save(tweet *model.Tweet) error {
	r.tweets = append(r.tweets, tweet)
	return nil
}

func (r *tweetRepository) Delete(id model.TweetID) error {
	for i, tweet := range r.tweets {
		if tweet.ID != id {
			continue
		}

		r.tweets = append(r.tweets[:i], r.tweets[i+1:]...)
		n := make([]*model.Tweet, len(r.tweets))
		copy(n, r.tweets)
		r.tweets = n
	}

	return nil
}
