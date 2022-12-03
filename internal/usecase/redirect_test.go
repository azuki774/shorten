package usecase

import (
	"azuki774/shorten/internal/model"
	"context"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestRedirectService_GetTargetURL(t *testing.T) {
	type fields struct {
		Logger       *zap.Logger
		DBRepository DBRepository
		HostName     string
	}
	type args struct {
		ctx      context.Context
		shortKey string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo model.URLShortInfoResponse
		wantErr  bool
	}{
		{
			name: "ok",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				HostName:     "http://localhost/",
			},
			args: args{
				ctx:      context.Background(),
				shortKey: "source",
			},
			wantInfo: model.URLShortInfoResponse{
				ShortURL:  "http://localhost/source",
				TargetURL: "http://localhost:8080/",
				ExpiredAt: time.Date(2000, 1, 23, 0, 0, 0, 0, jst),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedirectService{
				Logger:       tt.fields.Logger,
				DBRepository: tt.fields.DBRepository,
				HostName:     tt.fields.HostName,
			}
			gotInfo, err := r.GetTargetURL(tt.args.ctx, tt.args.shortKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedirectService.GetTargetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("RedirectService.GetTargetURL() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
