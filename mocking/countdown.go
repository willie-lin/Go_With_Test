package mocking

import (
	"fmt"
	"io"
)

const findWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, findWord)
}
