package redirector

import (
	"azuki774/shorten/internal/usecase"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (s *Server) middlewareLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			s.Logger.Info("access", zap.String("url", r.URL.Path), zap.String("X-Forwarded-For", r.Header.Get("X-Forwarded-For")))
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is the root page.\n")
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := mux.Vars(r)
	source := pathParam["source"]
	ctx := context.Background()
	info, err := s.RedirectService.GetTargetURL(ctx, source)
	if err != nil {
		if errors.Is(err, usecase.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, info.Target, http.StatusFound)
}
