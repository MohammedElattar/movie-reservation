// Package en
package en

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerValidation(b *i18.Bundle) {
	b.Register(i18.EnLocale, i18.Validation, map[string]string{
		"required": "field is required",
		"unique":   "must be unique",
	})
}
