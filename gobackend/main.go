package main

import (
    "log"
    "net/http"
    "github.com/pritesh/gobackend/routes"
    "github.com/pritesh/gobackend/config"
)

func main() {
	config.Connect()
    r := routes.SetupRouter()
    log.Println("Server running on port 8080...")
	log.Println("Available routes configured")
    log.Fatal(http.ListenAndServe(":8080", r))
}
