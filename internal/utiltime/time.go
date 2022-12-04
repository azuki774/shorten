package utiltime

import "time"

var NowFunc func() time.Time

func Now() time.Time {
	if NowFunc != nil {
		return NowFunc()
	}
	return time.Now()
}
