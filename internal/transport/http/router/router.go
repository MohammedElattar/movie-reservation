// Package router
package router

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/handlers"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	cfg *config.Config,
	jsonResponse *handlers.JsonResponse,
	userHandler *handlers.UserHandler,
) http.Handler {
	router := httprouter.New()

	// Health
	router.GET("/health", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok\n"))
	})

	// Auth Routes
	router.POST(
		"/auth/register",
		middleware.NewPipeline(userHandler.Register, cfg, jsonResponse).
			Through(middleware.GlobalMiddlewares(nil)...).
			Return(),
	)

	return router
}
