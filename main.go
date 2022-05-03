package main

import (
	"github.com/willie-lin/Go_With_Test/mocking"
	"os"
	"time"
)

func main() {
	//dependencyinjection.Greet(os.Stdout, "Elodie")

	//http.ListenAndServe(":5000", http.HandlerFunc(dependencyinjection.MyGreeterHandler))

	sleeper := &mocking.ConfigurableSleeper{1 * time.Second}
	mocking.Countdown(os.Stdout, sleeper)
}
