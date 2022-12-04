package usecase

import "azuki774/shorten/internal/model"

func newURLShortInfoResponse(u *model.URLShortInfo, hostname string) (info model.URLShortInfoResponse) {
	return model.URLShortInfoResponse{
		ShortURL:  hostname + u.ShortKey, // add hostname
		TargetURL: u.TargetURL,
		ExpiredAt: u.ExpiredAt,
	}
}
