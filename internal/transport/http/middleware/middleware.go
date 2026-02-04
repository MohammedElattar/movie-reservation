// Package middleware
package middleware

import (
	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/julienschmidt/httprouter"
)

type (
	Middleware func(next httprouter.Handle, cfg *config.Config) httprouter.Handle
)

var GlobalMiddlewares = []Middleware{
	Locale,
	ValidatePostSize,
}
