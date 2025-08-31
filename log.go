package middleware

import (
	"log"
	"net/http"
	"time"
)

type logWritter struct {
	W      http.ResponseWriter
	bytes  int
	status int
}

func (lw *logWritter) Header() http.Header {
	return lw.W.Header()
}

func (lw *logWritter) Write(body []byte) (int, error) {
	n, err := lw.W.Write(body)

	lw.bytes += n

	return n, err
}

func (lw *logWritter) WriteHeader(statusCode int) {
	lw.status = statusCode
	lw.W.WriteHeader(statusCode)
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &logWritter{W: w}

		start := time.Now()
		next.ServeHTTP(lw, r)
		dur := time.Since(start)

		// Common log line
		log.Printf("%s %s %d %s %s bytes=%d",
			r.Method,
			r.URL.Path,
			lw.status,
			r.RemoteAddr,
			dur.Round(time.Millisecond),
			lw.bytes,
		)
	})
}
