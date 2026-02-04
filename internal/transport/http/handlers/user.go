// Package handlers
package handlers

import (
	"fmt"
	"net/http"

	"github.com/MohammedElattar/movie-reservation/internal/domain/user"
	portsLogger "github.com/MohammedElattar/movie-reservation/internal/ports/logger"
	"github.com/MohammedElattar/movie-reservation/internal/transport/http/locale"
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	loginService *user.LoginService
	log          portsLogger.Logger
	i18          *i18.Bundle
}

func NewUserHandler(
	loginService *user.LoginService,
	log portsLogger.Logger,
	i18 *i18.Bundle,
) *UserHandler {
	return &UserHandler{
		loginService: loginService,
		log:          log,
		i18:          i18,
	}
}

func (h *UserHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	word := h.i18.Success(locale.FromContext(r.Context()), "name", "created")

	fmt.Println(word)
}
