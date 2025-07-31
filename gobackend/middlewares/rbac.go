package middlewares

import (
    "net/http"
    "github.com/pritesh/gobackend/utils"
    "github.com/gorilla/mux"
)


func RequireRoles(roles ...string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            user, ok := r.Context().Value("user").(*utils.JWTClaim)
            if !ok {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            for _, role := range roles {
                if user.Role == role {
                    next.ServeHTTP(w, r)
                    return
                }
            }

            http.Error(w, "Forbidden - insufficient role", http.StatusForbidden)
        })
    }
}

