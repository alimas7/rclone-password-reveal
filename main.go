package main

import (
	"fmt"
	"os"

	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run main.go <password_ofuscada>")
		os.Exit(1)
	}

	// Obtiene la contraseña ofuscada desde los argumentos
	obscuredPassword := os.Args[1]

	// Descifra la contraseña
	plainPassword, err := obscure.Reveal(obscuredPassword)
	if err != nil {
		fmt.Println("Error al descifrar la contraseña:", err)
		os.Exit(1)
	}

	fmt.Println("Contraseña original:", plainPassword)
}
