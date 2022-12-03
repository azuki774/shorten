package usecase

import (
	"azuki774/shorten/internal/model"
	"azuki774/shorten/internal/utiltime"
	"context"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
)

var l *zap.Logger

func init() {
	l, _ = zap.NewProduction()
}

func TestNewURLShortInfoFromReq(t *testing.T) {
	type args struct {
		req *model.URLRegistRequest
	}
	tests := []struct {
		name    string
		args    args
		wantU   model.URLShortInfo
		t       func() time.Time
		wantErr bool
	}{
		{
			name: "no expired",
			args: args{&model.URLRegistRequest{
				TargetURL: "target",
				ExpiredIn: 0,
			}},
			wantU: model.URLShortInfo{
				TargetURL: "target",
				ExpiredAt: time.Date(2038, 1, 1, 0, 0, 0, 0, jst),
			},
			t:       func() time.Time { return time.Date(2000, 1, 23, 0, 0, 0, 0, jst) },
			wantErr: false,
		},
		{
			name: "1h",
			args: args{&model.URLRegistRequest{
				TargetURL: "target",
				ExpiredIn: 3600,
			}},
			wantU: model.URLShortInfo{
				TargetURL: "target",
				ExpiredAt: time.Date(2000, 1, 23, 1, 0, 0, 0, jst),
			},
			t:       func() time.Time { return time.Date(2000, 1, 23, 0, 0, 0, 0, jst) },
			wantErr: false,
		},
		{
			name: "invalid args",
			args: args{&model.URLRegistRequest{
				TargetURL: "target",
				ExpiredIn: -1,
			}},
			wantU:   model.URLShortInfo{},
			t:       func() time.Time { return time.Date(2000, 1, 23, 0, 0, 0, 0, jst) },
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utiltime.NowFunc = tt.t
			gotU, err := NewURLShortInfoFromReq(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewURLShortInfoFromReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("NewURLShortInfoFromReq() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestRegistService_Regist(t *testing.T) {
	type fields struct {
		Logger       *zap.Logger
		DBRepository DBRepository
		KeyGenerator KeyGenerator
		HostName     string
	}
	type args struct {
		ctx context.Context
		req *model.URLRegistRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo model.URLShortInfoResponse
		t        func() time.Time
		wantErr  bool
	}{
		{
			name: "ok",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				KeyGenerator: &mockKeyGenerator{},
				HostName:     "http://localhost/",
			},
			args: args{
				ctx: context.Background(),
				req: &model.URLRegistRequest{
					TargetURL: "http://example.com/",
					ExpiredIn: 3600,
				},
			},
			wantInfo: model.URLShortInfoResponse{
				ShortURL:  "http://localhost/aaaaaaa",
				TargetURL: "http://example.com/",
				ExpiredAt: time.Date(2000, 1, 23, 1, 0, 0, 0, jst),
			},
			t:       func() time.Time { return time.Date(2000, 1, 23, 0, 0, 0, 0, jst) },
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utiltime.NowFunc = tt.t
			r := &RegistService{
				Logger:       tt.fields.Logger,
				DBRepository: tt.fields.DBRepository,
				KeyGenerator: tt.fields.KeyGenerator,
				HostName:     tt.fields.HostName,
			}
			gotInfo, err := r.Regist(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegistService.Regist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("RegistService.Regist() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
