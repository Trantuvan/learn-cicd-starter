package auth

import (
	"net/http"
	"testing"
)

type testCase struct {
	name   string
	auth   string
	apiKey string
	msg    string
}

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	cases := []testCase{
		{name: "no_authorization", auth: "", msg: "no authorization header included"},
		{name: "malformed_no_prefix_ApiKey", auth: "gibbresh fasf", msg: "malformed authorization header"},
		{name: "malformed_Apikey_len", auth: "Apikey", msg: "malformed authorization header"},
		{name: "happy_path", auth: "ApiKey 124", apiKey: "12"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			headers.Set("Authorization", tc.auth)
			got, err := GetAPIKey(headers)

			if err != nil {
				if err.Error() != tc.msg {
					t.Errorf("wanted an error: %s but didn't get one", err)
				}
			} else {
				if got != tc.apiKey {
					t.Errorf("got %s want %s", got, tc.apiKey)
				}
			}
		})
	}
}
