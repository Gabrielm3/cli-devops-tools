package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(tamanho int) string {
	rand.Seed(time.Now().UnixNano())
	if tamanho < 8 {
		fmt.Println("Tamanho mínimo de 8 caracteres")
	}
	b := make([]byte, tamanho)
	if _, err := rand.Read(b); err != nil {
		fmt.Println("Erro ao gerar a senha")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
