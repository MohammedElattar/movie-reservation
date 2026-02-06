package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ValidatePostSize(next httprouter.Handle, ctx *MiddlewareContext) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.ContentLength > ctx.Cfg.MaxPostSize {
			ctx.JsonResponse.ErrorResponse(
				r.Context(),
				w,
				ctx.JsonResponse.I18.Word(LocaleFromContext(r.Context()), "post_size_too_large"),
				nil,
				http.StatusRequestEntityTooLarge,
			)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, ctx.Cfg.MaxPostSize)
		next(w, r, ps)
	}
}
