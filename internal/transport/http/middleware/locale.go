package middleware

import (
	"net/http"
	"strings"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/locale"
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
	"github.com/julienschmidt/httprouter"
)

func Locale(next httprouter.Handle, cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		lang := detectLocale(r, i18.Locale(cfg.App.FallbackLocale))

		ctx := locale.WithLocale(r.Context(), lang)

		r = r.WithContext(ctx)

		next(w, r, ps)
	}
}

func detectLocale(r *http.Request, fallbackLocale i18.Locale) i18.Locale {
	if h := r.Header.Get("Accept-Language"); h != "" {
		return parseAcceptLanguage(h, fallbackLocale)
	}

	return fallbackLocale
}

func normalize(lang string) i18.Locale {
	return i18.Locale(strings.ToLower(strings.Split(lang, "-")[0]))
}

func parseAcceptLanguage(h string, fallbackLocale i18.Locale) i18.Locale {
	parts := strings.Split(h, ",")

	if len(parts) == 0 {
		return fallbackLocale
	}

	return normalize(parts[0])
}
