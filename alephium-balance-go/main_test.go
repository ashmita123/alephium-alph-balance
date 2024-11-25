package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetBalanceValidAddress(t *testing.T) {
    // Mock server that returns a valid balance
    mockBalance := "123.45"
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `{"balance": "%s"}`, mockBalance)
    }))
    defer server.Close()

    // Call the refactored getBalance with the mock server's URL
    address := "1H5pLrNnJCPh6snKkp8qodh2Tv3nJ2bYX5zD1AvStZAUG"
    balance, err := getBalance(address, server.URL)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if balance != 123.45 {
        t.Errorf("Expected balance 123.45, got %f", balance)
    }
}

func TestGetBalanceInvalidAddress(t *testing.T) {
    // Mock server that returns an error for invalid address
    mockError := "Invalid address"
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, `{"error": "%s"}`, mockError)
    }))
    defer server.Close()

    // Call the refactored getBalance with the mock server's URL
    address := "invalid_address"
    _, err := getBalance(address, server.URL)
    if err == nil {
        t.Error("Expected an error for invalid address, got nil")
    }
    expectedErrMsg := fmt.Sprintf("API Error: %s", mockError)
    if err.Error() != expectedErrMsg {
        t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
    }
}
