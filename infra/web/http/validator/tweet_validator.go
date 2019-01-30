package validator

import "net/http"

type ComposeTweetRequest struct {
	Content string
}

func ValidateToComposeTweet(r *http.Request) (*ComposeTweetRequest, error) {
	content := r.FormValue("content")
	if content == "" {
		return nil, newEmptyError("content")
	}

	return &ComposeTweetRequest{
		Content: content,
	}, nil
}
