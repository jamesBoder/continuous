

package auth

import (
	"testing"
	"net/http"
)



// create unit test for GetAPIKey. the function GetAPIKey returns a string and an error. 
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid api key",
			header:  "ApiKey my-secret-key",
			want:	"my-secret-key",
			wantErr: false,
		},
		{
			name:    "missing api key",
			header:  "ApiKey",
			want:	"",
			wantErr: true,
		},
		{
			name:    "malformed header",
			header:  "my-secret-key",
			want:	"",
			wantErr: true,
		},
		{
			name:    "no auth header",
			header:  "",
			want:	"",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.Header{}
			if tt.header != "" {
				//  expected format for that header value is usually just "ApiKey <key>"
				h.Set("Authorization", tt.header)
			}
			got, err := GetAPIKey(h)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}	

