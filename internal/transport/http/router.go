// Package router
package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userHandler *UserHandler,
	mwctx *MiddlewareContext,
) http.Handler {
	router := httprouter.New()

	// Health
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		mwctx.JsonResponse.OkResponse(r.Context(), w, struct{ Ok bool }{Ok: true}, nil)
	})

	// Auth Routes
	router.POST(
		"/auth/register",
		NewMiddlewarePipeline(userHandler.Register, mwctx).
			Through(GlobalMiddlewares(nil)...).
			Return(),
	)

	return router
}
