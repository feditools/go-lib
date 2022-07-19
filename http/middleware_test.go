package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const TestHeaderPermissionsPolicy = "Permissions-Policy"
const TestExpectedPermissionsPolicy = "interest-cohort=()"

func TestBlockFloc(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissionsPolicy := w.Header().Get(TestHeaderPermissionsPolicy)
		if permissionsPolicy != TestExpectedPermissionsPolicy {
			t.Errorf("%s header was incorrect, got: %v, want: %v.", TestHeaderPermissionsPolicy, permissionsPolicy, TestExpectedPermissionsPolicy)
		}
	})

	req := httptest.NewRequest("GET", "http://localhost", nil)
	blockFloc := BlockFloc(testHandler)
	blockFloc.ServeHTTP(httptest.NewRecorder(), req)
}

func TestBlockMissingUserAgent(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	req.Header.Set("User-Agent", "test")
	blockFloc := BlockMissingUserAgent(testHandler)
	blockFloc.ServeHTTP(resp, req)

	if resp.Result().StatusCode != 200 {
		t.Errorf("incorrect status code, got: %d, want: %d.", resp.Result().StatusCode, 200)
	}
}

func TestBlockMissingUserAgent_MissingUserAgent(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	blockFloc := BlockMissingUserAgent(testHandler)
	blockFloc.ServeHTTP(resp, req)

	if resp.Result().StatusCode != 400 {
		t.Errorf("incorrect status code, got: %d, want: %d.", resp.Result().StatusCode, 400)
	}
}
