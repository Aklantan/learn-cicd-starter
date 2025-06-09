package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			want:    "my-secret-key",
			wantErr: nil,
		},
		{
			name:    "Missing Authorization Header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "Malformed Authorization Header",
			headers: http.Header{"Authorization": []string{"Bearer token"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "Incomplete API Key Header",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}
