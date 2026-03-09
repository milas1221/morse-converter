package handlers

import (
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusInternalServerError)
        return
    }
    

    file, header, err := r.FormFile("myFile")
    if err != nil {
        http.Error(w, "Error getting file", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file", http.StatusInternalServerError)
        return
    }
    
    converted, err := service.DetectAndConvert(string(data))
    if err != nil {
        http.Error(w, "Error converting: " + err.Error(), http.StatusInternalServerError)
        return
    }
    
    ext := filepath.Ext(header.Filename)
    if ext == "" {
        ext = ".txt"
    }
    
    newFilename := time.Now().UTC().Format("20060102-150405") + ext
    
    w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Write([]byte(converted))
}