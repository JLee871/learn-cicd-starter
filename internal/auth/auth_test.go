package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{headers: http.Header{}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Malformed",
			args:    args{headers: http.Header{"Authorization": []string{"Api 1234"}}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Normal",
			args:    args{headers: http.Header{"Authorization": []string{"ApiKey 1234"}}},
			want:    "1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
