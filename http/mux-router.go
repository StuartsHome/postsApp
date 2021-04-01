package Router

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter
}
