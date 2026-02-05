package middleware

import (
	"net/http"
	"sync"
	"time"

	httputil "github.com/MohammedElattar/movie-reservation/internal/transport/http"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/context"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/locale"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

type client struct {
	LastSeen time.Time
	Limiter  *rate.Limiter
}

func RateLimiter(r rate.Limit, b int) Middleware {
	var (
		mu           sync.Mutex
		rateLimiters = make(map[string]*client)
	)

	return func(next httprouter.Handle, ctx *context.MiddlewareContext) httprouter.Handle {
		return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			ip := httputil.GetIP(req)

			mu.Lock()
			if _, ok := rateLimiters[ip]; !ok {
				rateLimiters[ip] = &client{
					Limiter: rate.NewLimiter(r, b),
				}
			}

			rateLimiters[ip].LastSeen = time.Now()
			mu.Unlock()

			if !rateLimiters[ip].Limiter.Allow() {
				ctx.JsonResponse.ErrorResponse(
					req.Context(),
					w,
					ctx.JsonResponse.I18.Word(locale.FromContext(req.Context()), "too_many_requests"),
					nil,
					http.StatusTooManyRequests,
				)
				return
			}

			next(w, req, ps)
		}
	}
}
