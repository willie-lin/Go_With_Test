package mocking

import (
	"fmt"
	"io"
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

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
