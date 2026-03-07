package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)


func DetectAndConvert(input string) (string, error) {

    input = strings.TrimSpace(input)
    

    if input == "" {
        return input, nil
    }
    

    if isMorseCode(input) {

        return morse.ToText(input), nil
    }
    

    return morse.ToMorse(input), nil
}


func isMorseCode(s string) bool {

    trimmed := strings.ReplaceAll(s, " ", "")
    

    if trimmed == "" {
        return false
    }
    

    for _, char := range trimmed {
        if char != '.' && char != '-' {
            return false
        }
    }
    return true
}