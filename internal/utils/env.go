package utils

import (
	"fmt"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("env PORT kosong, default ke " + port)
	}

	return ":" + port
}
