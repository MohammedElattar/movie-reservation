// Package ar
package ar

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerValidation(b *i18.Bundle) {
	b.Register(i18.ArLocale, i18.Validation, map[string]string{
		"required": "حقل %s مطلوب",
		"unique":   "حقل %s يجب ان يكون فريدا",
	})
}
