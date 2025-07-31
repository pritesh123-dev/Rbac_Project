package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/pritesh/gobackend/models"
    "github.com/pritesh/gobackend/config"
    "github.com/pritesh/gobackend/utils"
    "gorm.io/gorm"
)

var DB *gorm.DB // Assume initialized in main or config

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    hashedPwd, _ := utils.HashPassword(user.Password)
    user.Password = hashedPwd

    // user.Role = "viewer" // Default role
	config.DB.Create(&user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if config.DB == nil {
    http.Error(w, "Database not initialized", http.StatusInternalServerError)
    return
}
    var input models.User
    json.NewDecoder(r.Body).Decode(&input)

    var dbUser models.User
    result := config.DB.Where("username = ?", input.Username).First(&dbUser)

    if result.Error != nil || !utils.CheckPasswordHash(input.Password, dbUser.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateJWT(dbUser.Username, dbUser.Role)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
