package utils

import (
	"fmt"
	"log"
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

func GetSecret() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatalln("SECRET tidak boleh kosong")
	}
	return secret
}
