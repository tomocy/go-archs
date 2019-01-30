package request

import "github.com/tomocy/archs/domain/model"

type ComposeTweetRequest struct {
	UserID  model.UserID
	Content string
}

func NewComposeTweetRequest(userID model.UserID, content string) *ComposeTweetRequest {
	return &ComposeTweetRequest{
		UserID:  userID,
		Content: content,
	}
}
