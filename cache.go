// Package cache provides a `Cache` function
// that add `Cache-Control` and `Expires` headers to `http.Handler`s
package cache

import (
	"net/http"
	"strconv"
	"time"
)

// Cache add `Cache-Control` and `Expires` headers to the specified `http.Handler`
func Cache(days int, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if days < 1 {
			days = 1
		}

		age := days * 24 * 60 * 60 * 1000
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(age))
		t := time.Now().Add(time.Duration(time.Hour * 24 * time.Duration(days)))
		w.Header().Set("Expires", t.Format(time.RFC1123Z))

		next.ServeHTTP(w, r)
	})
}
