// Package locale
package locale

import (
	"context"

	"github.com/MohammedElattar/movie-reservation/pkg/i18"
)

type keyType struct{}

var key keyType = struct{}{}

func WithLocale(ctx context.Context, locale i18.Locale) context.Context {
	return context.WithValue(ctx, key, locale)
}

func FromContext(ctx context.Context) i18.Locale {
	if v, ok := ctx.Value(key).(i18.Locale); ok && v != "" {
		return v
	}

	return i18.EnLocale
}
