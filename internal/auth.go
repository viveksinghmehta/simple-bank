package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("secret-key-for-simple-bank")

func AuthenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the request headers
	token := c.GetHeader("Authorization")
	if token == "" {
		fmt.Println("Token missing in request")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is missing in the requst headers",
		})
		return
	}

	authToken := strings.Split(token, " ")[1]

	if authToken == "" {
		fmt.Println("Token missing in request")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is missing in the requst headers",
		})
		return
	}

	// Verify the token
	_, error := verifyToken(authToken)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Auth token is not valid",
		})
		return
	}
	c.Set("token", authToken)
	c.Next()
}

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
func verifyToken(tokenString string) (*jwt.Token, error) {
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
