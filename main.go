package main

import (
	"log"
	"net/http"

	"strconv"

	_ "github.com/reddyalready/otsp-srv/db"
)

func main() {
	port := 8080
	log.Printf("Starting server on %d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), NewRouter())
}
