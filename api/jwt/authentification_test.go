package jwt

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestToken struct {
	Method         string
	Url            string
	ExpectedStatus int
	User           string
	Password       string
}

func TestGetToken(t *testing.T) {
	var url = "http://0.0.0.0:8001/token/"
	getTokenBadUser := TestToken{"POST", url, http.StatusUnauthorized, "badUser", "dskfop"}
	getTokenValid := TestToken{"POST", url, http.StatusOK, "admin", "password"}

	tests := []TestToken{getTokenBadUser, getTokenValid}

	for _, test := range tests {
		body := strings.NewReader(`user=` + test.User + `&password=` + test.Password)
		req, err := http.NewRequest("POST", "http://0.0.0.0:8001/token", body)
		if err != nil {
			t.Fatalf("Fail to create get token request")
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()
		GetToken(recorder, req)

		if recorder.Code != test.ExpectedStatus {
			t.Fatalf("FAIL test, expected status %d", test.ExpectedStatus)
		}
	}
}
