package middlewares

import (
    "net/http"
    "github.com/pritesh/gobackend/utils"
    "github.com/gorilla/mux"
)


func RequireRole(role string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            user, ok := r.Context().Value("user").(*utils.JWTClaim)
            if !ok || user.Role != role {
                http.Error(w, "Forbidden - insufficient role", http.StatusForbidden)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
