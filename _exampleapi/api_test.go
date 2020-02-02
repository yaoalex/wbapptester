package exampleapi_test

import "testing"

func TestGetProducts(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{name: "valid test case"},
		{name: "invalid test case"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func TestGetProduct(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{name: "valid test case"},
		{name: "invalid test case"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func TestPurchaseProduct(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{name: "valid test case"},
		{name: "invalid test case"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
