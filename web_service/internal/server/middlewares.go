package server

import (
	"context"
	"net/http"
	"time"
)

func (srv *server) contextExpire(h http.HandlerFunc) http.HandlerFunc {
	srv.logger.Info("contextExpire")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Minute)
		defer cancel()

		r = r.WithContext(ctx)
		h(w, r)
	}
}
