package middleware

import (
	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/julienschmidt/httprouter"
)

type Pipeline struct {
	handler httprouter.Handle
	cfg     *config.Config
}


func NewPipeline(handler httprouter.Handle, cfg *config.Config) *Pipeline {
	return &Pipeline{cfg: cfg, handler: handler}
}

func (p *Pipeline) Through(mws ...Middleware) *Pipeline {
	for i := len(mws) - 1; i >= 0; i-- {
		p.handler = mws[i](p.handler, p.cfg)
	}

	return p
}

func (p *Pipeline) Return() httprouter.Handle {
	return httprouter.Handle(p.handler)
}

// CallStack, args
