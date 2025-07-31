package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/pritesh/gobackend/controllers"
    "github.com/pritesh/gobackend/middlewares"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Auth routes
    r.HandleFunc("/register", controllers.Register).Methods("POST")
    r.HandleFunc("/login", controllers.Login).Methods("POST")

    // Protected routes
  	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTAuth)

	// 👇 /upload route for admin and editor
	api.Handle("/upload", middlewares.RequireRoles("admin", "editor")(http.HandlerFunc(controllers.UploadDocument))).Methods("POST")

	// 👇 View and Download routes for all roles
	api.Handle("/documents", middlewares.RequireRoles("admin", "editor", "viewer")(http.HandlerFunc(controllers.GetDocuments))).Methods("GET")
	api.Handle("/documents/{id}", middlewares.RequireRoles("admin", "editor", "viewer")(http.HandlerFunc(controllers.DownloadDocument))).Methods("GET")
	api.Handle("/users", middlewares.RequireRoles("admin")(http.HandlerFunc(controllers.GetUsers))).Methods("GET")
	api.Handle("/users/{id}", middlewares.RequireRoles("admin")(http.HandlerFunc(controllers.DeleteUser))).Methods("DELETE")
	api.Handle("/users/{id}/role", middlewares.RequireRoles("admin")(http.HandlerFunc(controllers.ChangeUserRole))).Methods("PATCH")

    return r
}
