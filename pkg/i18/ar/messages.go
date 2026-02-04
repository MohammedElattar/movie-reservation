package ar

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerMessages(b *i18.Bundle) {
	b.Register(i18.ArLocale, i18.Messages, map[string]string{
		"created":             "تم إنشاء %s بنجاح",
		"name":                "الإسم",
		"resource_created":    "تم الإنشاء بنجاح",
		"post_size_too_large": "حجم الطلب كبير جدًا. يرجى تقليل حجم البيانات أو الملفات المرفقة.",
		"too_many_requests": "لقد أرسلت عددًا كبيرًا من الطلبات خلال فترة قصيرة. يرجى الانتظار ثم المحاولة مرة أخرى لاحقًا.",
	})
}
