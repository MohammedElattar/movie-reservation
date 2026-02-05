// Package context
package context

import (
	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/ports/logger"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/httpresponse"
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
)

type MiddlewareContext struct {
	I18 *i18.Bundle
	Log logger.Logger
	JsonResponse *httpresponse.JsonResponse
	Cfg *config.Config
}

func NewContext(i18 *i18.Bundle, log logger.Logger, jsonResponse *httpresponse.JsonResponse, cfg *config.Config) *MiddlewareContext{
	return &MiddlewareContext{
		I18: i18,
		Log: log,
		JsonResponse: jsonResponse,
		Cfg: cfg,
	}
}
