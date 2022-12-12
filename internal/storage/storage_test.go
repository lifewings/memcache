package storage

import (
	"context"
	"testing"
)

func TestStorage_Delete(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name       string
		argsSet    args
		argsDelete args
		want       string
		wantErr    bool
	}{
		{
			name: "Success delete",
			argsSet: args{
				key:   "test1",
				value: "success",
			},
			argsDelete: args{
				key:   "test1",
				value: "success",
			},
			want:    DeleteSuccess,
			wantErr: false,
		},
		{
			name: "Error delete",
			argsSet: args{
				key:   "test1",
				value: "success",
			},
			argsDelete: args{
				key:   "test2",
				value: "success",
			},
			want:    "",
			wantErr: true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			_, err := s.Set(ctx, tt.argsSet.key, tt.argsSet.value)
			if err != nil {
				t.Errorf("Set for delete %v", err)
				return
			}
			got, err := s.Delete(ctx, tt.argsDelete.key)
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

func TestStorage_Get(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		argsSet args
		argsGet args
		want    string
		wantErr bool
	}{
		{
			name: "success get",
			argsSet: args{
				key:   "test1",
				value: "success",
			},
			argsGet: args{
				key:   "test1",
				value: "success",
			},
			want:    "success",
			wantErr: false,
		},
		{
			name: "error get",
			argsSet: args{
				key:   "test1",
				value: "success",
			},
			argsGet: args{
				key:   "test2",
				value: "success",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error get with delete",
			argsSet: args{
				key:   "test1",
				value: "success",
			},
			argsGet: args{
				key:   "test2",
				value: "success",
			},
			want:    "",
			wantErr: true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			_, err := s.Set(ctx, tt.argsSet.key, tt.argsSet.value)
			if err != nil {
				t.Errorf("Set for delete %v", err)
				return
			}
			got, err := s.Get(ctx, tt.argsGet.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success set",
			args: args{
				key:   "test1",
				value: "success",
			},
			want:    "",
			wantErr: false,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			got, err := s.Set(ctx, tt.args.key, tt.args.value)
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
