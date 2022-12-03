package usecase

import (
	"azuki774/shorten/internal/model"
	"azuki774/shorten/internal/utiltime"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

var jst *time.Location

func init() {
	j, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	jst = j
}

var ErrRecordNotFound = errors.New("not found")

type DBRepository interface {
	GetTargetURL(ctx context.Context, shortKey string, t time.Time) (u model.URLShortInfo, err error)
	RegistDB(ctx context.Context, info model.URLShortInfo) (err error)
}
type RedirectService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	HostName     string
}

func (r *RedirectService) GetTargetURL(ctx context.Context, shortKey string) (info model.URLShortInfoResponse, err error) {
	u, err := r.DBRepository.GetTargetURL(ctx, shortKey, utiltime.Now())
	if err != nil && !errors.Is(err, ErrRecordNotFound) {
		r.Logger.Error("failed to get target URL", zap.Error(err))
		return model.URLShortInfoResponse{}, err
	} else if errors.Is(err, ErrRecordNotFound) {
		r.Logger.Warn("not found target URL", zap.Error(err))
		return model.URLShortInfoResponse{}, err
	}

	// Add HostName
	info = newURLShortInfoResponse(&u, r.HostName)
	return info, nil
}
