package main

import (
	"log"
	"net/http"

	"github.com/phati/circleci-demo/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", server.InitRouter()))
}