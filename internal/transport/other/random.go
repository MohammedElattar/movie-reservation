package random

import (
	"context"

	"github.com/MohammedElattar/movie-reservation/internal/transport/http"
)


func NewAdapter() {
	http.LocaleFromContext(context.Background())
}
