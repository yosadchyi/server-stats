package main

import (
	"awesomeProject2/pkg/stats"
	"log"
	"net/http"
)

func main() {
	addresses := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}

	err := http.ListenAndServe(":8080", stats.HandleServersStatsRequest(addresses, stats.FetchLoadOverHTTP))
	if err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
