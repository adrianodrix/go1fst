package utils

import (
	"fmt"
	"time"
)

// GetCurrentTime retorna a hora atual formatada
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FormatMessage formata uma mensagem com timestamp
func FormatMessage(message string) string {
	return fmt.Sprintf("[%s] %s", GetCurrentTime(), message)
}

// função privada (só pode ser usada dentro do pacote utils)
func internalHelper() string {
	return "função interna"
}

// Função pública (pode ser usada por outros pacotes)
func GetInternalInfo() string {
	return internalHelper() + " - chamada externamente"
}
