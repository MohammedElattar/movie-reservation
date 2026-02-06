package http

import (
	"net/http"
	"sync"
	"time"

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

	return func(next httprouter.Handle, ctx *MiddlewareContext) httprouter.Handle {
		return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			ip := GetIP(req)

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
					ctx.JsonResponse.I18.Word(LocaleFromContext(req.Context()), "too_many_requests"),
					nil,
					http.StatusTooManyRequests,
				)
				return
			}

			next(w, req, ps)
		}
	}
}
