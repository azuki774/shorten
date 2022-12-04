package repository

import (
	"azuki774/shorten/internal/model"
	"azuki774/shorten/internal/usecase"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type DBRepository struct {
	Conn *gorm.DB
}

func (d *DBRepository) CloseDB() (err error) {
	dbconn, err := d.Conn.DB()
	if err != nil {
		return err
	}
	return dbconn.Close()
}

func (d *DBRepository) GetTargetURL(ctx context.Context, source string, t time.Time) (u model.URLShortInfo, err error) {
	var r model.URLTable
	err = d.Conn.Where("short_key = ?", source).Where("expired_at > ?", t).First(&r).Error
	if err != nil {
		return model.URLShortInfo{}, err
	}

	u = model.NewURLShortInfo(&r)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = usecase.ErrRecordNotFound
		return model.URLShortInfo{}, err
	} else if err != nil {
		return model.URLShortInfo{}, err
	}
	return u, nil
}

func (d *DBRepository) RegistDB(ctx context.Context, info model.URLShortInfo) (err error) {
	r := model.NewURLTable(&info)
	err = d.Conn.Create(&r).Error
	if err != nil {
		return err
	}
	return nil
}
