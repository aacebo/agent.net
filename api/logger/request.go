package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func Request(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lw := &loggingResponseWriter{w, http.StatusOK}

			next.ServeHTTP(lw, r)

			if r.Method == http.MethodOptions {
				return
			}

			status := Text(strconv.Itoa(lw.statusCode)).Bold()

			if lw.statusCode < 400 {
				status = status.GreenForeground()
			} else if lw.statusCode < 500 {
				status = status.YellowForeground()
			} else {
				status = status.RedForeground()
			}

			elapse := time.Now().Sub(start)
			elapseText := Text(fmt.Sprintf("%dms", elapse.Milliseconds()))

			if elapse < 500*time.Millisecond {
				elapseText = elapseText.GreenForeground()
			} else if elapse < 5*time.Second {
				elapseText = elapseText.YellowForeground()
			} else {
				elapseText = elapseText.RedForeground()
			}

			log.Info(fmt.Sprintf(
				"%s %s %s %s",
				Text(r.Method).YellowForeground().Bold().String(),
				status.String(),
				Text(r.URL.Path).Underline().String(),
				elapseText.String(),
			))
		})
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (self *loggingResponseWriter) WriteHeader(code int) {
	self.statusCode = code
	self.ResponseWriter.WriteHeader(code)
}
