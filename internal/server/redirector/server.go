package redirector

import (
	"azuki774/shorten/internal/model"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type RedirectService interface {
	GetTargetURL(ctx context.Context, key string) (info model.URLShortInfo, err error)
}

type Server struct {
	Logger          *zap.Logger
	RedirectService RedirectService
}

func (s *Server) Start(ctx context.Context) error {
	s.Logger.Info("redirector start")
	router := mux.NewRouter()
	s.addRecordFunc(router)

	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	ctxIn, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var errCh = make(chan error)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	<-ctxIn.Done()
	if nerr := server.Shutdown(ctx); nerr != nil {
		s.Logger.Error("failed to shutdown server", zap.Error(nerr))
		return nerr
	}

	err := <-errCh
	if err != nil && err != http.ErrServerClosed {
		s.Logger.Error("failed to close server", zap.Error(err))
		return err
	}

	s.Logger.Info("http server close gracefully")
	return nil
}

func (s *Server) addRecordFunc(r *mux.Router) {
	r.HandleFunc("/", s.rootHandler)
	r.HandleFunc("/{source}", s.getHandler)
	r.Use(s.middlewareLogging)
}
