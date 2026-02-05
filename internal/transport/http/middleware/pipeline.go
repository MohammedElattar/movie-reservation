package middleware

import (
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/context"
	"github.com/julienschmidt/httprouter"
)

type Pipeline struct {
	ctx *context.MiddlewareContext
	handler httprouter.Handle
}


func NewPipeline(handler httprouter.Handle, ctx *context.MiddlewareContext) *Pipeline {
	return &Pipeline{ctx: ctx, handler: handler}
}

func (p *Pipeline) Through(mws ...Middleware) *Pipeline {
	for i := len(mws) - 1; i >= 0; i-- {
		p.handler = mws[i](p.handler, p.ctx)
	}

	return p
}

func (p *Pipeline) Return() httprouter.Handle {
	return httprouter.Handle(p.handler)
}

// CallStack, args
