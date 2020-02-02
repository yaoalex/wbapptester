package exampleapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// THIS IS GENERATED CODE BY WEBAPPTESTER
// you will need to edit this code to suit your API's needs

func TestGetProducts (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
	}{
		{
			Name: "GetProducts: valid test case",
			ExpectedStatus: http.StatusOK,
		},
		{
			Name: "GetProducts: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetProducts)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
func TestGetProduct (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
	}{
		{
			Name: "GetProduct: valid test case",
			ExpectedStatus: http.StatusOK,
		},
		{
			Name: "GetProduct: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetProduct)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
func TestPurchaseProduct (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
	}{
		{
			Name: "PurchaseProduct: valid test case",
			ExpectedStatus: http.StatusOK,
		},
		{
			Name: "PurchaseProduct: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PurchaseProduct)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
