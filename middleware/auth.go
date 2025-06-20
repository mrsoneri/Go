package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-auth-api/auth"
	"github.com/dgrijalva/jwt-go"
)
// Claims struct that will be encoded to JWT
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateToken creates a JWT token for a specific username
func GenerateToken(username string) (string, error) {
    // Set token expiration time
    expirationTime := time.Now().Add(24 * time.Hour)
    
    // Create the claims
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    
    // Create the token using the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Sign the token with the secret key
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
            c.Abort()
            return
        }

        token, err := auth.ValidateToken(tokenString)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse claims"})
            c.Abort()
            return
        }

        username := claims["username"].(string)
        c.Set("username", username)
        c.Next()
    }
}