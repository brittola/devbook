package main

import (
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Create()

	fmt.Println("Servidor rodando na porta 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
