package middleware

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/handlers"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/locale"
	"github.com/julienschmidt/httprouter"
)

func ValidatePostSize(next httprouter.Handle, cfg *config.Config, jsonResponse *handlers.JsonResponse) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.ContentLength > cfg.MaxPostSize {
			jsonResponse.ErrorResponse(
				r.Context(),
				w,
				jsonResponse.I18.Word(locale.FromContext(r.Context()), "post_size_too_large"),
				nil,
				http.StatusRequestEntityTooLarge,
			)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, cfg.MaxPostSize)
		next(w, r, ps)
	}
}
