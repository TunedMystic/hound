package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tunedmystic/hound/app/server"
)

func main() {
	server := server.NewServer()

	addr := "0.0.0.0:8000"
	fmt.Printf("Running server on %v ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
