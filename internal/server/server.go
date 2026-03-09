// package server

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
// )

// type Server struct {
//     logger     *log.Logger
//     httpServer *http.Server
// }

// func NewServer(logger *log.Logger) *Server {
//     router := http.NewServeMux()
    
//     router.HandleFunc("/", handlers.IndexHandler)
//     router.HandleFunc("/upload", handlers.UploadHandler)
    
//     httpServer := &http.Server{
//         Addr:         ":8080",
//         Handler:      router,
//         ErrorLog:     logger,
//         ReadTimeout:  5 * time.Second,
//         WriteTimeout: 10 * time.Second,
//         IdleTimeout:  15 * time.Second,
//     }
    
//     return &Server{
//         logger:     logger,
//         httpServer: httpServer,
//     }
// }

// // Run запускает сервер
// func (s *Server) Run() error {
//     return s.httpServer.ListenAndServe()
// }