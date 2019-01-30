package handler

type Handler interface {
	AuthenticationHandler
	TweetHandler
	UserHandler
}

type handler struct {
	AuthenticationHandler
	TweetHandler
	UserHandler
}

func NewHandler(
	authHandler AuthenticationHandler,
	tweetHandler TweetHandler,
	userHandler UserHandler,
) Handler {
	return &handler{
		AuthenticationHandler: authHandler,
		TweetHandler:          tweetHandler,
		UserHandler:           userHandler,
	}
}
