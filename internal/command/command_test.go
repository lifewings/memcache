package command

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"memcache/internal/command/mocks"
)

func TestServer_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		err     error
	}{
		{
			name: "cmd success delete",
			args: args{
				key: "test",
			},
			want:    "test",
			wantErr: false,
			err:     nil,
		},
		{
			name: "cmd error delete",
			args: args{
				key: "test",
			},
			want:    "",
			wantErr: true,
			err:     errors.New("has error"),
		},
	}
	ctx := context.Background()
	storage := mocks.Storager{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&storage, zap.NewNop())
			storage.On("Delete", mock.Anything, mock.Anything).Return(tt.args.key, tt.err).Once()
			got, err := s.Delete(ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Get(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   bool
		wantErr bool
		err     error
	}{
		{
			name: "cmd success get",
			args: args{
				key:   "test1",
				value: "success",
			},
			want:    "success",
			want1:   true,
			wantErr: false,
			err:     nil,
		},
		{
			name: "cmd error key get",
			args: args{
				key:   "test1",
				value: "",
			},
			want:    "",
			want1:   false,
			wantErr: false,
			err:     ErrorNotFound,
		},
		{
			name: "cmd error get",
			args: args{
				key:   "test1",
				value: "",
			},
			want:    "",
			want1:   false,
			wantErr: true,
			err:     errors.New("error cache"),
		},
	}
	ctx := context.Background()
	storage := mocks.Storager{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&storage, zap.NewNop())
			storage.On("Get", mock.Anything, mock.Anything).Return(tt.args.value, tt.err).Once()
			got, got1, err := s.Get(ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestServer_Set(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		err     error
	}{
		{
			name: "cmd success set",
			args: args{
				key:  "test1",
				data: "success",
			},
			want:    "success",
			wantErr: false,
			err:     nil,
		},
	}
	ctx := context.Background()
	storage := mocks.Storager{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(&storage, zap.NewNop())
			storage.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.data, tt.err).Once()
			got, err := s.Set(ctx, tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}
