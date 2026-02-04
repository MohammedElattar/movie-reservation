package ar

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerMessages(b *i18.Bundle) {
	b.Register(i18.ArLocale, i18.Messages, map[string]string{
		"created": "تم إنشاء %s بنجاح",
		"name":    "الإسم",
	})
}
