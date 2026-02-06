package http

import (
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/domain/user"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	loginService *user.LoginService
	mwctx *MiddlewareContext
}

func NewUserHandler(
	loginService *user.LoginService,
	mwctx *MiddlewareContext,
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
	word := h.mwctx.I18.Success(LocaleFromContext(r.Context()), "name", "created")

	h.mwctx.JsonResponse.CreatedResponse(r.Context(), w, struct {
		Message string
	}{
		Message: word,
	})
}
