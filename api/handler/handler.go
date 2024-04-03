package handler

type handlerV1 struct{}

// NewHandlerV1 is a constructor for handlerV1
type HandlerV1Config struct{}

func NewHandlerV1(h *HandlerV1Config) *handlerV1 {
	return &handlerV1{}
}
