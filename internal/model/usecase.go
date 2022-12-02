package model

import "time"

type URLShortInfo struct {
	Source    string    `json:"source"`
	Target    string    `json:"target"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewURLShortInfo(d *URLTable) (u URLShortInfo) {
	return URLShortInfo{
		Source:    d.Source,
		Target:    d.Target,
		ExpiredAt: d.ExpiredAt,
	}
}
