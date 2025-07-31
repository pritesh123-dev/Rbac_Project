package controllers


import (
    "fmt"
    "io"
    "net/http"
    "os"
	"encoding/json"
    "path/filepath"
    "time"
	"github.com/gorilla/mux"
    "github.com/pritesh/gobackend/config"
    "github.com/pritesh/gobackend/models"
    "github.com/pritesh/gobackend/utils"
)

func UploadDocument(w http.ResponseWriter, r *http.Request) {
    // Parse the multipart form
    err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
    if err != nil {
        http.Error(w, "File too big", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("file")
if err != nil {
    http.Error(w, "Error retrieving the file", http.StatusBadRequest)
    return
}
defer file.Close()

f, err := os.Create("uploads/" + handler.Filename)
if err != nil {
    http.Error(w, "Error saving file", http.StatusInternalServerError)
    return
}
defer f.Close()

_, err = io.Copy(f, file)
if err != nil {
    http.Error(w, "Error saving file", http.StatusInternalServerError)
    return
}
    defer file.Close()

    // Get user from context
    user := r.Context().Value("user").(*utils.JWTClaim)

    // Generate unique filename
    fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
    savePath := filepath.Join("uploads", fileName)
	

    // Create destination file
    dst, err := os.Create(savePath)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    io.Copy(dst, file)

    // Save file metadata to DB
    document := models.Document{
        Filename:  handler.Filename,
        Filepath:      savePath,
        UploadedBy: user.Username,
    }

    config.DB.Create(&document)

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "File uploaded successfully")
}


func GetDocuments(w http.ResponseWriter, r *http.Request) {
    var docs []models.Document
    config.DB.Find(&docs)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(docs)
}



func DownloadDocument(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var doc models.Document
    result := config.DB.First(&doc, id)
    if result.Error != nil {
        http.Error(w, "Document not found", http.StatusNotFound)
        return
    }

    http.ServeFile(w, r, doc.Filepath)
}
