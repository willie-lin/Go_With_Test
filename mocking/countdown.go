package mocking

import (
	"fmt"
	"io"
	"time"
)

const findWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {

	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, findWord)
}
