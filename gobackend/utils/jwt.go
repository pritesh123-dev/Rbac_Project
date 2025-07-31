package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtKey = []byte("your-secret-key")

type JWTClaim struct {
    Username string
    Role     string
    jwt.RegisteredClaims
}

func GenerateJWT(username, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &JWTClaim{
        Username: username,
        Role:     role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaim)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}
