package model

import (
	"fmt"
	"time"
)

type URLShortInfo struct {
	ShortKey  string    `json:"short_key"`
	TargetURL string    `json:"target_url"`
	ExpiredAt time.Time `json:"expired_at"`
}

type URLShortInfoResponse struct {
	ShortURL  string    `json:"short_url"`
	TargetURL string    `json:"target_url"`
	ExpiredAt time.Time `json:"expired_at"`
}

type URLRegistRequest struct {
	TargetURL string `json:"target_url"`
	ExpiredIn int    `json:"expired_in"`
}

func NewURLShortInfo(d *URLTable) (u URLShortInfo) {
	return URLShortInfo{
		ShortKey:  d.ShortKey,
		TargetURL: d.TargetURL,
		ExpiredAt: d.ExpiredAt,
	}
}

func NewURLTable(d *URLShortInfo) (u URLTable) {
	return URLTable{
		ShortKey:  d.ShortKey,
		TargetURL: d.TargetURL,
		ExpiredAt: d.ExpiredAt,
	}
}

func ValidateRegistRequest(req *URLRegistRequest) error {
	if req.ExpiredIn < 0 {
		return fmt.Errorf("invalid args")
	}
	return nil
}
