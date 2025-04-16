package auth_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKeyTable(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		header  http.Header
		wantErr error
	}{
		"no Authorization header": {
			header:  make(http.Header),
			wantErr: auth.ErrNoAuthHeaderIncluded,
		},
		"no ApiKey value": {
			header:  http.Header{"Authorization": []string{"ApiKey"}},
			wantErr: auth.ErrMalformedAuthHeader,
		},
		"invalid Authorization method": {
			header:  http.Header{"Authorization": []string{"Basic", "user:pass"}},
			wantErr: auth.ErrMalformedAuthHeader,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := auth.GetAPIKey(tc.header)
			if err == nil {
				t.Fatalf("GetAPIKey(%+v) expected to get an error", tc.header)
			}

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("GetAPIKey(%+v) got error %v, want %v", tc.header, err, tc.wantErr)
			}
		})
	}
}
