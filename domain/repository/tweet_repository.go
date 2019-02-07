package repository

import "github.com/tomocy/archs/domain/model"

type TweetRepository interface {
	NextID() model.TweetID
	Save(tweet *model.Tweet) error
	Delete(id model.TweetID) error
}
