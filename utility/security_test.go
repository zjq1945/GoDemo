package utility

import "testing"
import "net/http"

func TestAuthenticateRequest(t *testing.T) {
	req, err := http.NewRequest("Get", "/Consumer/GetAll", nil)
	if err != nil {
		t.Fatal(err)
	}
	if AuthenticateRequest(req.Header) {
		t.Error("Expected false")
	}
	req, err = http.NewRequest("Get", "/Consumer/GetAll", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authentication", "hasToken")
	if !AuthenticateRequest(req.Header) {
		t.Error("Expected true")
	}

}
