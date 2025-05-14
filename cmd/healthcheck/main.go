package main

import (
	"fmt"
	"log"
	"net/http"

	"weather-api/cmd/config"
)

func main() {
	_, err := http.Get(fmt.Sprintf("http://%s:%d/-/healthy", config.Host, config.Port))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
