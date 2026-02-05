// Package handlers
package handlers

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/domain/user"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/context"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/locale"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	loginService *user.LoginService
	mwctx *context.MiddlewareContext
}

func NewUserHandler(
	loginService *user.LoginService,
	mwctx *context.MiddlewareContext,
) *UserHandler {
	return &UserHandler{
		loginService: loginService,
		mwctx: mwctx,
	}
}

func (h *UserHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	word := h.mwctx.I18.Success(locale.FromContext(r.Context()), "name", "created")

	h.mwctx.JsonResponse.CreatedResponse(r.Context(), w, struct {
		Message string
	}{
		Message: word,
	})
}
