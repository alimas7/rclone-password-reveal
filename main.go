package main

import (
	"fmt"
	"os"

	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <obscured_password>")
		os.Exit(1)
	}

	// Get the obscured password from the arguments
	obscuredPassword := os.Args[1]

	// Reveal the original password
	plainPassword, err := obscure.Reveal(obscuredPassword)
	if err != nil {
		fmt.Println("Error revealing the password:", err)
		os.Exit(1)
	}

	fmt.Println("Original password:", plainPassword)
}
