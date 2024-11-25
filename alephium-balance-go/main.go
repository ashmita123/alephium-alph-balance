package main

import (
    "encoding/json"
    "errors"
    "flag"
    "fmt"
    "net/http"
)

// AddressResponse represents the API response structure.
type AddressResponse struct {
    Balance string `json:"balance"`
    Error   string `json:"error,omitempty"`
}

// getBalance fetches the ALPH balance of a given Alephium address.
func getBalance(address string, baseURL string) (float64, error) {
    url := fmt.Sprintf("%s/addresses/%s", baseURL, address)
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var addrResp AddressResponse
    if err := json.NewDecoder(resp.Body).Decode(&addrResp); err != nil {
        return 0, err
    }

    if resp.StatusCode != http.StatusOK {
        if addrResp.Error != "" {
            return 0, fmt.Errorf("API Error: %s", addrResp.Error)
        }
        return 0, fmt.Errorf("error fetching balance: %s", resp.Status)
    }

    if addrResp.Balance == "" {
        return 0, errors.New("invalid response format")
    }

    var balance float64
    _, err = fmt.Sscanf(addrResp.Balance, "%f", &balance)
    if err != nil {
        return 0, errors.New("balance is not a valid number")
    }

    return balance, nil
}

func main() {
    addressPtr := flag.String("address", "", "Alephium address")
    flag.Parse()

    if *addressPtr == "" {
        fmt.Println("Please provide an Alephium address using the -address flag.")
        return
    }

    // Use the mainnet API endpoint
    baseURL := "https://backend.mainnet.alephium.org"

    balance, err := getBalance(*addressPtr, baseURL)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Balance of %s: %f ALPH\n", *addressPtr, balance)
}
