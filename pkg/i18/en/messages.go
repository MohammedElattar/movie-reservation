package en

import "github.com/MohammedElattar/movie-reservation/pkg/i18"

func registerMessages(b *i18.Bundle) {
	b.Register(i18.EnLocale, i18.Messages, map[string]string{
		"created":             "%s has been created succesfully",
		"name":                "Name",
		"resource_created":    "Resource has been created successfully",
		"data_fetched":        "Data fetched successfully",
		"post_size_too_large": "The request is too large. Reduce the payload or file size.",
		"too_many_requests":   "You have sent too many requests in a short period. Please wait and try again later.",
		"success_operation": "Operation has been made successfully",
	})
}
