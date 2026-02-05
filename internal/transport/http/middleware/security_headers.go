package middleware

import (
	"net/http"
	"os"

	"github.com/MohammedElattar/movie-reservation/internal/transport/http/context"
	"github.com/julienschmidt/httprouter"
)

func AddSecurityHeaders(next httprouter.Handle, ctx *context.MiddlewareContext) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Prevent clickjacking
		w.Header().Set("X-Frame-Options", "DENY")

		// Prevent MIME sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		if os.Getenv("APP_ENV") == "production" {
			// HSTS - Force HTTPS
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		// Referrer policy
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		next(w, r, ps)
	}
}
