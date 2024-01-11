package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()

		logger := log.New(logFile, "", 0)

		var details string
		if r.Method == "POST" {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Printf("Error reading request body: %v", err)
			} else {
				details = string(bodyBytes)
			}
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		wrappedWriter := NewResponseWriter(w)
		next.ServeHTTP(wrappedWriter, r)

		timestamp := time.Now().Format("02-01-2006 15:04:05")
		if r.Method == "POST" {
			logger.Printf("[%s] %s %s %d %s", timestamp, r.Method, r.RequestURI, wrappedWriter.Status(), details)
		} else {
			logger.Printf("[%s] %s %s %d", timestamp, r.Method, r.RequestURI, wrappedWriter.Status())
		}
	})
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) Status() int {
	return rw.statusCode
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
	if code >= 400 {
		log.Printf("Error: Status code %d", code)
	}
}
