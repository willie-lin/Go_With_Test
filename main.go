package main

import (
	"github.com/willie-lin/Go_With_Test/mocking"
	"os"
)

func main() {
	//dependencyinjection.Greet(os.Stdout, "Elodie")

	//http.ListenAndServe(":5000", http.HandlerFunc(dependencyinjection.MyGreeterHandler))

	mocking.Countdown(os.Stdout)
}
