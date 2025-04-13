package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)


func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
			input 	http.Header
			wantKey string
			wantErr error
	}{
			"simple": {
				input: http.Header{"Authorization": []string{"ApiKey aKey"}},
				wantKey: "aKey",
				wantErr: nil,
			},
			"empty": {
				input: http.Header{"Authorization": []string{}},
				wantKey: "",
				wantErr: ErrNoAuthHeaderIncluded,
			},
			"malformed": {
				input: http.Header{"Authorization": []string{"ApiKey"}},
				wantKey: "",
				wantErr: ErrMalformedAuthHeader,
			},
	}

	for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				key, err := GetAPIKey(tc.input)

				keyDiff := cmp.Diff(tc.wantKey, key)
				if keyDiff != "" {
					t.Fatal(keyDiff)
				}

				errDiff := cmp.Diff(tc.wantErr, err, cmpopts.EquateErrors())
				if errDiff != "" {
					t.Fatal(errDiff)
				}
			})
	}
}