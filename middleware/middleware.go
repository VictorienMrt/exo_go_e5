package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// LoggerMiddleware is a middleware function that logs HTTP requests.
// It creates or opens a log file and records details of each request.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Open or create the log file.
		logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()

		// Create a new logger instance.
		logger := log.New(logFile, "", 0)

		// Read and log the request body for POST requests.
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

		// Wrap the response writer for status code logging.
		wrappedWriter := NewResponseWriter(w)
		next.ServeHTTP(wrappedWriter, r)

		// Log the request details.
		timestamp := time.Now().Format("02-01-2006 15:04:05")
		if r.Method == "POST" {
			logger.Printf("[%s] %s %s %d %s", timestamp, r.Method, r.RequestURI, wrappedWriter.Status(), details)
		} else {
			logger.Printf("[%s] %s %s %d", timestamp, r.Method, r.RequestURI, wrappedWriter.Status())
		}
	})
}

// NewResponseWriter creates a new responseWriter instance.
// It is used to capture and log the HTTP status code of responses.
func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

// responseWriter is a wrapper around http.ResponseWriter that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Status returns the HTTP status code.
func (rw *responseWriter) Status() int {
	return rw.statusCode
}

// WriteHeader captures and logs the HTTP status code,
// and then calls the underlying ResponseWriter's WriteHeader method.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
	if code >= 400 {
		log.Printf("Error: Status code %d", code)
	}
}
