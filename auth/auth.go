package auth

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var secretKey = []byte("your_secret_key") // Change this to a more secure key in production

// GenerateToken generates a new JWT token for a given username.
func GenerateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration set to 1 hour
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// ValidateToken validates a JWT token and returns the claims.
func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
}