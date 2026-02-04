package middleware

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/julienschmidt/httprouter"
)

func ValidatePostSize(next httprouter.Handle, cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.ContentLength > cfg.MaxPostSize {
			http.Error(w, "post_size_too_large", http.StatusRequestEntityTooLarge)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, cfg.MaxPostSize)
		next(w, r, ps)
	}
}
