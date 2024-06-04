package main

import (
	"devbook-api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.Create()

	log.Fatal(http.ListenAndServe(":3000", r))
}
