package exampleapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// THIS IS GENERATED CODE BY WBAPPTESTER
// you will need to edit this code to suit your needs

func TestGetProducts(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		ExpectedStatus int
		MuxVars        map[string]string
	}{
		{
			Name:           "GetProducts: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars: map[string]string{
				"show_empty": "valid_value",
			},
		},
		{
			Name:           "GetProducts: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars: map[string]string{
				"show_empty": "invalid_value",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetProducts)

			req = mux.SetURLVars(req, tc.MuxVars)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
func TestGetProduct(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		ExpectedStatus int
		MuxVars        map[string]string
	}{
		{
			Name:           "GetProduct: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars: map[string]string{
				"title": "valid_value",
			},
		},
		{
			Name:           "GetProduct: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars: map[string]string{
				"title": "invalid_value",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetProduct)

			req = mux.SetURLVars(req, tc.MuxVars)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
func TestGetTest(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		ExpectedStatus int
		MuxVars        map[string]string
	}{
		{
			Name:           "GetTest: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars:        map[string]string{},
		},
		{
			Name:           "GetTest: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars:        map[string]string{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetTest)

			req = mux.SetURLVars(req, tc.MuxVars)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
func TestPurchaseProduct(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		ExpectedStatus int
		MuxVars        map[string]string
	}{
		{
			Name:           "PurchaseProduct: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars: map[string]string{
				"title": "valid_value",
			},
		},
		{
			Name:           "PurchaseProduct: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars: map[string]string{
				"title": "invalid_value",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PurchaseProduct)

			req = mux.SetURLVars(req, tc.MuxVars)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
