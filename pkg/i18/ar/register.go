package ar

import (
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
)

func Register(b *i18.Bundle) {
	registerMessages(b)
	registerValidation(b)
}
