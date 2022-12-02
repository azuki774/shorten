package factory

import (
	"azuki774/shorten/internal/server/redirector"
	"azuki774/shorten/internal/usecase"
	"azuki774/shorten/repository"
	"fmt"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, err := config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		fmt.Printf("failed to create logger: %v\n", err)
	}
	return l, err
}

func NewRedirectService(l *zap.Logger, db *repository.DBRepository) (ap *usecase.RedirectService) {
	return &usecase.RedirectService{Logger: l, DBRepository: db}
}

func NewRedirecter(l *zap.Logger, ap *usecase.RedirectService) *redirector.Server {
	return &redirector.Server{Logger: l, RedirectService: ap}
}
