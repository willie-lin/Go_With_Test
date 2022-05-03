package main

import (
	"github.com/willie-lin/Go_With_Test/mocking"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func main() {
	//dependencyinjection.Greet(os.Stdout, "Elodie")

	//http.ListenAndServe(":5000", http.HandlerFunc(dependencyinjection.MyGreeterHandler))

	sleeper := &ConfigurableSleeper{1 * time.Second}
	mocking.Countdown(os.Stdout, sleeper)
}
