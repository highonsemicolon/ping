package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Jwks struct {
	Keys []Jwk `json:"keys"`
}

type Jwk struct {
	Kid string   `json:"kid"`
	X5c []string `json:"x5c"`
}

// GetPemCert fetches the Auth0 JWKS and retrieves the PEM certificate for token verification
func GetPemCert(domain string) (string, error) {
	resp, err := http.Get("https://" + domain + "/.well-known/jwks.json")
	if err != nil {
		return "", fmt.Errorf("failed to fetch JWKS: %v", err)
	}
	defer resp.Body.Close()

	var jwks Jwks
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read JWKS response body: %v", err)
	}
	err = json.Unmarshal(body, &jwks)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JWKS: %v", err)
	}

	if len(jwks.Keys) == 0 || len(jwks.Keys[0].X5c) == 0 {
		return "", errors.New("no certificates found in JWKS")
	}

	cert := fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", jwks.Keys[0].X5c[0])
	return cert, nil
}
