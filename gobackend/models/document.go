package models

import ("gorm.io/gorm"
// "time"
)

type Document struct {
    gorm.Model
    Filename   string
    Filepath   string
    UploadedBy string
}
