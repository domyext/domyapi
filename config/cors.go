package domyApi

import (
	"log"  
	"net/http"
)

// Daftar origins yang diizinkan
var Origins = []string{
	"https://naskah.bukupedia.co.id",
	"https://chatgpl.do.my.id",
	"https://do.my.id",
	"https://in.my.id",
	"https://my.my.id",
	"https://whatsauth.github.io",
	"https://www.do.my.id",
	"https://domyext.github.io",
	"https://domy.buxxed.me",
}

// Fungsi untuk memeriksa apakah origin diizinkan
func isAllowedOrigin(origin string) bool {
	for _, o := range Origins {
		if o == origin {
			return true
		}
	}
	log.Printf("Origin %s not allowed", origin)
	return false
}

// Fungsi untuk mengatur header CORS
func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")

	if isAllowedOrigin(origin) {
		// Set CORS headers for the preflight request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Login")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return true
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		return false
	}

	return false
}
