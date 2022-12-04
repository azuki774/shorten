package factory

import (
	"azuki774/shorten/internal/server/redirector"
	"azuki774/shorten/internal/server/register"
	"azuki774/shorten/internal/usecase"
	"azuki774/shorten/repository"
	"fmt"
	"os"

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
	hn := os.Getenv("HOST_NAME")
	return &usecase.RedirectService{Logger: l, DBRepository: db, HostName: hn}
}

func NewRedirecter(l *zap.Logger, ap *usecase.RedirectService) *redirector.Server {
	return &redirector.Server{Logger: l, RedirectService: ap}
}

func NewRegistService(l *zap.Logger, db *repository.DBRepository, kg *repository.SourceGenerator) (ap *usecase.RegistService) {
	hn := os.Getenv("HOST_NAME")
	return &usecase.RegistService{Logger: l, DBRepository: db, KeyGenerator: kg, HostName: hn}
}

func NewRegister(l *zap.Logger, ap *usecase.RegistService) *register.Server {
	return &register.Server{Logger: l, RegistService: ap}
}

func NewSourceGenerator() *repository.SourceGenerator {
	return &repository.SourceGenerator{}
}
