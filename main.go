package main

import (
	"github.com/willie-lin/Go_With_Test/dependencyinjection"
	"net/http"
	"os"
)

func main() {
	dependencyinjection.Greet(os.Stdout, "Elodie")

	http.ListenAndServe(":5000", http.HandlerFunc(dependencyinjection.MyGreeterHandler))

}
