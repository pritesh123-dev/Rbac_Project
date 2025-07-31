package routes

import (
    // "net/http"


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
	api.Use(middlewares.RequireRole("admin"))

    api.HandleFunc("/upload", controllers.UploadDocument).Methods("POST")
    api.HandleFunc("/documents", controllers.GetDocuments).Methods("GET")
    api.HandleFunc("/documents/{id}", controllers.DownloadDocument).Methods("GET")

    return r
}
