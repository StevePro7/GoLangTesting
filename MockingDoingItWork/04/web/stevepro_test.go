package web

import (
	"net/http"
	"testing"
)

func TestGetFixedValue(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactiveAndReset()

	httpmock.RegisterResponder("GET", "https://example.com/fixedvalue",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Accept") != "application/json" {
				t.Errorf("expected Accep: application/json header, got: %s", req.Header.Get("Accept"))
			}
			resp, err := httpmock.NewJsonResponse(http.StatusOK, map[string]interface{} {
				"value": "fixed",
			})
			return resp, err
		},
	}

	value, _ := GetFixedValue("https://example.com")
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}
}