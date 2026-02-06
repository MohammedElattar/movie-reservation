package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MohammedElattar/movie-reservation/pkg/i18"
	"github.com/MohammedElattar/movie-reservation/pkg/json"
)

type HTTPData = any

var (
	EmptyObject = struct{}{}
	EmptySlice  = []struct{}{}
)

type (
	HTTPStatusCode int32
	HTTPResponse   struct {
		Data    HTTPData          `json:"data"`
		Message string            `json:"message"`
		Code    HTTPStatusCode    `json:"code"`
		Errors  map[string]string `json:"errors"`
	}
	JsonResponse struct {
		I18 *i18.Bundle
	}
)

func NewJsonResponseWriter(i18 *i18.Bundle) *JsonResponse {
	return &JsonResponse{
		I18: i18,
	}
}

func (res *JsonResponse) ErrorResponse(
	ctx context.Context,
	w http.ResponseWriter,
	message string,
	data any,
	code int,
) {
	streamedResponse(w, data, message, code, nil)
}

func (res *JsonResponse) CreatedResponse(ctx context.Context, w http.ResponseWriter, data any) {
	streamedResponse(
		w,
		data,
		res.I18.Word(LocaleFromContext(ctx), "resource_created"),
		http.StatusCreated,
		nil,
	)
}

func (res *JsonResponse) ResourceResponse(ctx context.Context, w http.ResponseWriter, data any) {
	streamedResponse(
		w,
		data,
		res.I18.Word(LocaleFromContext(ctx), "data_fetched"),
		http.StatusOK,
		nil,
	)
}

func (res *JsonResponse) OkResponse(ctx context.Context, w http.ResponseWriter, data any, message *string) {
	var msg string

	if message == nil {
		msg = res.I18.Word(LocaleFromContext(ctx), "success_operation")
	} else {
		msg = *message
	}

	streamedResponse(
		w,
		data,
		msg,
		http.StatusOK,
		nil,
	)
}

func (res JsonResponse) PaginatedResponse(ctx context.Context, w http.ResponseWriter, data any) {
	streamedResponse(
		w,
		data,
		res.I18.Word(LocaleFromContext(ctx), "data_fetched"),
		http.StatusOK,
		nil,
	)
}

func baseResponse(
	w http.ResponseWriter,
	data any,
	message string,
	code int,
	errors map[string]string,
) {
	response := HTTPResponse{
		Data:    data,
		Message: message,
		Code:    HTTPStatusCode(code),
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(code))

	buf, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in json encoding %v", err), http.StatusInternalServerError)
		return;
	}

	_, err = w.Write(buf)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when writing json %v", err), http.StatusInternalServerError)
		return
	}
}

func streamedResponse(
	w http.ResponseWriter,
	data any,
	message string,
	code int,
	errors map[string]string,
) {
	response := HTTPResponse{
		Data:    data,
		Message: message,
		Code:    HTTPStatusCode(code),
		Errors:  errors,
	}

	// Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(code)

	enc := json.NewEncoder(w)

	if err := enc.Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	// Flush if supported
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}
