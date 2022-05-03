package mocking

import (
	"fmt"
	"io"
	"time"
)

const findWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		//fmt.Fprintln(out, i)
	}
	for i := countdownStart; i > 0; i-- {
		//sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, findWord)
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
