package proxy

import "net/http"

func NewProxy(handle http.Handler, fn ...http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, f := range fn {
			f(w, r)
		}
		if handle != nil {
			handle.ServeHTTP(w, r)
		}
	})
}
