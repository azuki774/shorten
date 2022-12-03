package usecase

import (
	"azuki774/shorten/internal/model"
	"context"
	"time"
)

type mockDBRepository struct{}
type mockKeyGenerator struct{}

func (m *mockDBRepository) GetTargetURL(ctx context.Context, shortKey string, t time.Time) (u model.URLShortInfo, err error) {
	return model.URLShortInfo{
		ShortKey:  "source",
		TargetURL: "http://localhost:8080/",
		ExpiredAt: time.Date(2000, 1, 23, 0, 0, 0, 0, jst),
	}, nil
}

func (m *mockDBRepository) RegistDB(ctx context.Context, info model.URLShortInfo) (err error) {
	return nil
}

func (m *mockKeyGenerator) Generate(n int) (shortKey string, err error) {
	ret := ""
	for i := 0; i < n; i++ {
		ret += "a"
	}
	return ret, nil
}
