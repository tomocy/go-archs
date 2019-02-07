package response

import "github.com/tomocy/archs/domain/model"

type TweetUsecaseResponser interface {
	ResponseTweet(tweet *model.Tweet) *TweetResponse
}

type TweetResponse struct {
	ID      string
	UserID  string
	Content string
}

func NewTweetResponse(id, userID, content string) *TweetResponse {
	return &TweetResponse{
		ID:      id,
		UserID:  userID,
		Content: content,
	}
}
