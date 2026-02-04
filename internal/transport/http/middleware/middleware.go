// Package middleware
package middleware

import (
	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/handlers"
	"github.com/julienschmidt/httprouter"
)

type (
	Middleware func(next httprouter.Handle, cfg *config.Config, jsonResponse *handlers.JsonResponse) httprouter.Handle
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

	return append(globalMiddlewares, RateLimiter(1, 0))
}
