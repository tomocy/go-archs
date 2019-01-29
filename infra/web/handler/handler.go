package handler

type Handler interface {
	AuthenticationHandler
	UserHandler
}

type handler struct {
	AuthenticationHandler
	UserHandler
}

func NewHandler(
	authHandler AuthenticationHandler,
	userHandler UserHandler,
) Handler {
	return &handler{
		AuthenticationHandler: authHandler,
		UserHandler:           userHandler,
	}
}
