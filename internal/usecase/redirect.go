package usecase

import (
	"azuki774/shorten/internal/model"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

var ErrRecordNotFound = errors.New("not found")

type DBRepository interface {
	GetTargetURL(ctx context.Context, source string, t time.Time) (u model.URLShortInfo, err error)
}
type RedirectService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
}

func (r *RedirectService) GetTargetURL(ctx context.Context, source string) (info model.URLShortInfo, err error) {
	u, err := r.DBRepository.GetTargetURL(ctx, source, time.Now())
	if err != nil && !errors.Is(err, ErrRecordNotFound) {
		r.Logger.Error("failed to get target URL", zap.Error(err))
		return model.URLShortInfo{}, err
	} else if errors.Is(err, ErrRecordNotFound) {
		r.Logger.Warn("not found target URL", zap.Error(err))
		return model.URLShortInfo{}, err
	}
	return u, nil
}
