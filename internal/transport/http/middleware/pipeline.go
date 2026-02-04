package middleware

import (
	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/handlers"
	"github.com/julienschmidt/httprouter"
)

type Pipeline struct {
	handler httprouter.Handle
	cfg     *config.Config
	jsonResponse *handlers.JsonResponse
}


func NewPipeline(handler httprouter.Handle, cfg *config.Config, jsonResponse *handlers.JsonResponse) *Pipeline {
	return &Pipeline{cfg: cfg, handler: handler, jsonResponse: jsonResponse}
}

func (p *Pipeline) Through(mws ...Middleware) *Pipeline {
	for i := len(mws) - 1; i >= 0; i-- {
		p.handler = mws[i](p.handler, p.cfg, p.jsonResponse)
	}

	return p
}

func (p *Pipeline) Return() httprouter.Handle {
	return httprouter.Handle(p.handler)
}

// CallStack, args
