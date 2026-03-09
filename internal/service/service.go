// package service

// import (
// 	"errors"
// 	"strings"

// 	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
// )

// // DetectAndConvert определяет, текст это или морзе, и конвертирует в противоположный формат
// func DetectAndConvert(input string) (string, error) {
//     if input == "" {
//         return "", errors.New("empty input")
//     }
    
//     // Убираем пробелы в начале и конце
//     trimmed := strings.TrimSpace(input)
    
//     // Определяем, это морзе или текст
//     // Морзе содержит только точки, тире и пробелы
//     isMorse := true
//     for _, char := range trimmed {
//         if char != '.' && char != '-' && char != ' ' {
//             isMorse = false
//             break
//         }
//     }
    
//     // Конвертируем
//     if isMorse {
//         // Морзе -> текст
//         return morse.ToText(trimmed), nil
//     } else {
//         // Текст -> морзе
//         return morse.ToMorse(trimmed), nil
//     }
// }