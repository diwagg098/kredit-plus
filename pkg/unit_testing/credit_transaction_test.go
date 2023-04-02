package unit_testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCreditTransaction(t *testing.T) {
	s := []struct {
		name    string
		wantErr error
		content string
		auth    string
	}{
		{
			name:    "Success Case",
			wantErr: nil,
			content: ``,
		},
	}

	for _, cases := range s {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/credit-transaction/all", bytes.NewBuffer([]byte(cases.content)))
		req.Header.Set("Authorization", cases.auth)
		req.Header.Set("Content-Type", "application/json")

		t.Run(cases.name, func(t *testing.T) {
			ctrl.CreditTransactionList(httptest.NewRecorder(), req)
		})
	}
}

func TestGetAcademyDetail(t *testing.T) {

	s := []struct {
		name    string
		wantErr error
		content string
		auth    string
	}{
		{
			name:    "Success Case",
			wantErr: nil,
			content: ``,
		},
	}

	for _, cases := range s {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/credit-transaction/1", bytes.NewBuffer([]byte(cases.content)))
		req.Header.Set("Authorization", cases.auth)
		req.Header.Set("Content-Type", "application/json")

		t.Run(cases.name, func(t *testing.T) {
			ctrl.FindByIdCreditTransaction(httptest.NewRecorder(), req)
		})
	}
}
