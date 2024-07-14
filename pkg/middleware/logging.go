package middleware

import (
	"log"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile("logs/logstash.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening log file:", err)
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
