package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		log.Printf("========= ผลที่ได้้จากการเรียก API=======")
		log.Printf("========= Start=======")

		log.Printf(">>> %s %s | Body: %s", r.Method, r.URL.Path, string(bodyBytes))

		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Printf("<<< API: <<< Body: %s", string(bodyBytes))
		log.Printf("<<< API: <<< Status: %d", rw.status)

		log.Printf("=========End=======")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
