package api

import (
	"net/http"
	"regexp"
)

// Middleware is used to handle the CORS Preflight
type Middleware struct {
	OriginRule string
}

func (m *Middleware) allowedOrigin(origin string) bool {
	if m.OriginRule == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(m.OriginRule, origin); matched {
		return true
	}
	return false
}

// cors manage CORS preflight
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// in advance we need to change this for the appopiate OriginRule rule
			m := Middleware{OriginRule: "*"}
			if m.allowedOrigin(r.Header.Get("Origin")) {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			}
			return
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
