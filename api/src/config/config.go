package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnectionString é a string de conexão ao banco de dados
	ConnectionString = ""

	// Port representa a porta onde a API vai estar rodando
	Port = 0
)

// Load inicializa as variáveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 3000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
}
