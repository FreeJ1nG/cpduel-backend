package util

import (
	"fmt"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.status = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recoder := &statusRecorder{w, http.StatusOK}
		now := time.Now()
		defer func() {
			fmt.Printf(
				"method=%s, url=%s, host=%s, path=%s, duration=%s, status=%d\n",
				r.Method,
				r.RequestURI,
				r.Host,
				r.URL.Path,
				time.Since(now).String(),
				recoder.status,
			)
		}()
		next.ServeHTTP(recoder, r)
	})
}
