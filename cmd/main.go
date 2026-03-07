package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

    logger := log.New(os.Stdout, "MORSE: ", log.LstdFlags|log.Lshortfile)
    

    srv := server.NewServer(logger)
    

    logger.Println("Запуск сервиса конвертации азбуки Морзе")
    if err := srv.Run(); err != nil {
        logger.Fatal("Ошибка сервера:", err)
    }
}