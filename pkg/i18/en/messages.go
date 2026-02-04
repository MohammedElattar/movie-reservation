package en

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerMessages(b *i18.Bundle) {
	b.Register(i18.EnLocale, i18.Messages, map[string]string{
		"created": "%s has been created succesfully",
		"name":    "Name",
	})
}
