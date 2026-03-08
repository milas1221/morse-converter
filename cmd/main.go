package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
    logger := log.New(os.Stdout, "MORSE: ", log.LstdFlags)
    
    s := server.NewServer(logger)
    
    logger.Println("Сервер запущен на http://localhost:8080")
    if err := s.Run(); err != nil {
        logger.Fatal("Ошибка запуска сервера:", err)
    }
}