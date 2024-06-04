package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	r := router.Create()

	fmt.Println("Servidor rodando na porta 3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), r))
}
