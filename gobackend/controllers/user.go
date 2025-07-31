// UploadDocument godoc
// @Summary Upload a document
// @Tags Documents
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Document File"
// @Success 201 {string} string "File uploaded successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Router /upload [post]
// @Security BearerAuth

package controllers

import (
    "encoding/json"
    "net/http"
	"github.com/gorilla/mux"
    "github.com/pritesh/gobackend/config"
    "github.com/pritesh/gobackend/models"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    result := config.DB.Select("id", "username", "role", "created_at").Find(&users)
    if result.Error != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User
    result := config.DB.First(&user, id)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Hard delete using Unscoped
    config.DB.Unscoped().Delete(&user)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted permanently"})
}


func ChangeUserRole(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    var body struct {
        Role string `json:"role"`
    }
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Optional: only allow valid roles
    validRoles := map[string]bool{
        "admin": true, "editor": true, "viewer": true,
    }
    if !validRoles[body.Role] {
        http.Error(w, "Invalid role", http.StatusBadRequest)
        return
    }

    result := config.DB.Model(&models.User{}).Where("id = ?", userID).Update("role", body.Role)
    if result.Error != nil {
        http.Error(w, "Failed to update role", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": "User role updated successfully",
    })
}
