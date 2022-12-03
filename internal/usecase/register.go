package usecase

import (
	"azuki774/shorten/internal/model"
	"azuki774/shorten/internal/utiltime"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

var ErrInvalidArgs = errors.New("invalid args")

const srcNum = 7

type KeyGenerator interface {
	Generate(n int) (source string, err error)
}

type RegistService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	KeyGenerator KeyGenerator
	HostName     string
}

func NewURLShortInfoFromReq(req *model.URLRegistRequest) (u model.URLShortInfo, err error) {
	if err := model.ValidateRegistRequest(req); err != nil {
		return model.URLShortInfo{}, err
	}
	expiredAt := utiltime.Now().Add(time.Duration(req.ExpiredIn) * time.Second)
	if req.ExpiredIn == 0 {
		// expired_in = 0 -> expired_at : inf
		expiredAt = time.Date(2038, 1, 1, 0, 0, 0, 0, jst)
	}
	return model.URLShortInfo{
		TargetURL: req.TargetURL,
		ExpiredAt: expiredAt,
	}, nil

}

func (r *RegistService) Regist(ctx context.Context, req *model.URLRegistRequest) (info model.URLShortInfoResponse, err error) {
	u, err := NewURLShortInfoFromReq(req)
	if err != nil {
		r.Logger.Warn("invalid args", zap.Error(err))
		return model.URLShortInfoResponse{}, ErrInvalidArgs
	}

	// Generate source
	shortKey, err := r.KeyGenerator.Generate(srcNum)
	if err != nil {
		r.Logger.Error("failed to generate shortKey", zap.Error(err))
		return model.URLShortInfoResponse{}, err
	}
	u.ShortKey = shortKey

	err = r.DBRepository.RegistDB(ctx, u)
	if err != nil {
		r.Logger.Error("failed to insert to DB", zap.Error(err))
		return model.URLShortInfoResponse{}, err
	}

	// Add hostname
	info = newURLShortInfoResponse(&u, r.HostName)
	r.Logger.Info("regist new info", zap.String("short_key", shortKey), zap.String("target_url", req.TargetURL))
	return info, nil
}
