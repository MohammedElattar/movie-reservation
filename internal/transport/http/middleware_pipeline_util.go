package http

import (
	"github.com/julienschmidt/httprouter"
)

type Pipeline struct {
	ctx *MiddlewareContext
	handler httprouter.Handle
}


func NewMiddlewarePipeline(handler httprouter.Handle, ctx *MiddlewareContext) *Pipeline {
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
