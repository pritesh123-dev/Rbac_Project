package middlewares

import (
    "github.com/pritesh/gobackend/utils"
    "net/http"
    "strings"
    "context"
)

func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
        claims, err := utils.ValidateToken(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
