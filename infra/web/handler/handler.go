package handler

type Handler interface {
	UserHandler
}

type handler struct {
	UserHandler
}

func NewHandler(u UserHandler) Handler {
	return &handler{
		UserHandler: u,
	}
}
