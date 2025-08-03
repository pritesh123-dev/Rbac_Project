package main

import (
	"log"
    "net/http"
    "github.com/pritesh/gobackend/routes"
    "github.com/pritesh/gobackend/config"
    // "github.com/pritesh/gobackend/middlewares" 
    "github.com/rs/cors"
)

func main() {
	config.Connect()
    r := routes.SetupRouter()
    log.Println("Server running on port 8080...")
	log.Println("Available routes configured")
    // r.Use(middlewares.CORSMiddleware)
    handler := cors.Default().Handler(r)
    log.Fatal(http.ListenAndServe(":8080", handler))

    
}
