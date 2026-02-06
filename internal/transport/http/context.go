package http

import (
	"context"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/ports/logger"
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
)

// ----------------------------
// Locale Context
// ----------------------------

type localKeyType struct{}

var key localKeyType = localKeyType{}

func WithLocale(ctx context.Context, locale i18.Locale) context.Context {
	return context.WithValue(ctx, key, locale)
}

func LocaleFromContext(ctx context.Context) i18.Locale {
	if v, ok := ctx.Value(key).(i18.Locale); ok && v != "" {
		return v
	}

	return i18.EnLocale
}

// ----------------------------
// Middleware Context
// ----------------------------

type MiddlewareContext struct {
	I18          *i18.Bundle
	Log          logger.Logger
	JsonResponse *JsonResponse
	Cfg          *config.Config
}

func NewMiddlewareContext(i18 *i18.Bundle, log logger.Logger, jsonResponse *JsonResponse, cfg *config.Config) *MiddlewareContext {
	return &MiddlewareContext{
		I18:          i18,
		Log:          log,
		JsonResponse: jsonResponse,
		Cfg:          cfg,
	}
}
