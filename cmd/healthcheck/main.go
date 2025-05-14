package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("http://localhost:8000/-/healthy")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
