// Package http
package http

import (
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// Check X-Forwarded-For header (set by Nginx)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Get first IP (client IP)
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Fallback to X-Real-IP
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	// Fallback to RemoteAddr
	return r.RemoteAddr
}

// func IsSecure(r *http.Request) bool {
//     return r.Header.Get("X-Forwarded-Proto") == "https"
// }
