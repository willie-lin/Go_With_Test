package main

import (
	"github.com/willie-lin/Go_With_Test/app"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(app.PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
