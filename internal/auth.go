package internal

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("secret-key-for-simple-bank")

// Function to create JWT tokens with claims
func CreateToken(email string, expire int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                                                    // Subject (user identifier)
		"iss": "simple_bank_project",                                    // Issuer
		"iat": time.Now().Unix(),                                        // Issued at
		"ext": time.Now().Add(time.Hour * time.Duration(expire)).Unix(), // Expiration time
	})

	tokenString, error := claims.SignedString(secretkey)
	if error != nil {
		panic("Can't create auth token")
	}
	return tokenString, nil
}

// Function to verify JWT tokens
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
