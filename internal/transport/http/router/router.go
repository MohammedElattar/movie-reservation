// Package router
package router

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/transport/http/context"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/handlers"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userHandler *handlers.UserHandler,
	mwctx *context.MiddlewareContext,
) http.Handler {
	router := httprouter.New()

	// Health
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		mwctx.JsonResponse.OkResponse(r.Context(), w, struct{Ok bool}{Ok: true}, nil);
	})

	// Auth Routes
	router.POST(
		"/auth/register",
		middleware.NewPipeline(userHandler.Register, mwctx).
			Through(middleware.GlobalMiddlewares(nil)...).
			Return(),
	)

	return router
}
