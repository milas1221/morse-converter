package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {

    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, "Ошибка загрузки страницы: "+err.Error(), http.StatusInternalServerError)
        return
    }
    

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Ошибка отображения страницы: "+err.Error(), http.StatusInternalServerError)
        return
    }
}


func UploadHandler(logger *log.Logger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        err := r.ParseMultipartForm(10 << 20)
        if err != nil {
            logger.Printf("Ошибка парсинга формы: %v", err)
            http.Error(w, "Не удалось обработать форму", http.StatusInternalServerError)
            return
        }
        

        file, header, err := r.FormFile("file")
        if err != nil {
            logger.Printf("Ошибка получения файла: %v", err)
            http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
            return
        }
        defer file.Close()
        
        logger.Printf("Получен файл: %s", header.Filename)
        

        content, err := io.ReadAll(file)
        if err != nil {
            logger.Printf("Ошибка чтения файла: %v", err)
            http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
            return
        }
        

        result, err := service.DetectAndConvert(string(content))
        if err != nil {
            logger.Printf("Ошибка конвертации: %v", err)
            http.Error(w, "Не удалось конвертировать содержимое", http.StatusInternalServerError)
            return
        }

        ext := filepath.Ext(header.Filename)
        

        outputFilename := time.Now().UTC().Format("20060102-150405") + ext
        

        outputFile, err := os.Create(outputFilename)
        if err != nil {
            logger.Printf("Ошибка создания выходного файла: %v", err)
            http.Error(w, "Не удалось создать выходной файл", http.StatusInternalServerError)
            return
        }
        defer outputFile.Close()
        

        _, err = outputFile.WriteString(result)
        if err != nil {
            logger.Printf("Ошибка записи в выходной файл: %v", err)
            http.Error(w, "Не удалось записать в выходной файл", http.StatusInternalServerError)
            return
        }
        
        logger.Printf("Создан выходной файл: %s", outputFilename)
        

        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(result))
    }
}