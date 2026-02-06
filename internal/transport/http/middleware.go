// Package middleware
package http

import (
	"github.com/julienschmidt/httprouter"
)

type (
	Middleware func(next httprouter.Handle, ctx *MiddlewareContext) httprouter.Handle
)

var globalMiddlewares = []Middleware{
	AddSecurityHeaders,
	ValidatePostSize,
	Locale,
}

func GlobalMiddlewares(ignoreGlobalLimiter *struct{}) []Middleware {
	if ignoreGlobalLimiter != nil {
		return globalMiddlewares
	}

	return append(globalMiddlewares, RateLimiter(1, 60))
}
