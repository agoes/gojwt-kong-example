package token

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyEnvName = "JWT_PRIVATE_KEY"
	algorithm      = "RS256"

	// DefaultExpiresIn is the default value for token expiration in seconds
	DefaultExpiresIn = 3600
)

// CreateToken generate new JWT for specific issuer
func CreateToken(issuer string, expiresIn int) string {
	signBytes, err := getPrivateKey()
	fatal(err)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	t := jwt.New(jwt.GetSigningMethod(algorithm))
	t.Claims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
		NotBefore: time.Now().Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}

	token, err := t.SignedString(signKey)
	fatal(err)

	return token
}

func getPrivateKeyFromEnv() string {
	return os.Getenv(privKeyEnvName)
}

func getPrivateKey() ([]byte, error) {
	privateKey := os.Getenv(privKeyEnvName)
	if privateKey != "" {
		return []byte(privateKey), nil
	}

	return nil, fmt.Errorf("Missing or empty env variable: %s", privKeyEnvName)
}

func fatal(err error) {
	if err != nil {
		log.Fatalf("JWT error : %s", err)
	}
}
